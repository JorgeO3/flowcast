// mongotx_test.go
package mongotx

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/JorgeO3/flowcast/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// Asegúrate de que tu logger implementa correctamente logger.Interface con métodos Error y Info.
// Si no, ajusta las llamadas al logger en el paquete mongotx según lo explicado anteriormente.

// TestMongoTx_Run_Success verifica que una transacción exitosa no retorna errores.
func TestMongoTx_Run_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		// Configurar el mock para respuestas exitosas
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // StartTransaction
			mtest.CreateSuccessResponse(), // CommitTransaction
		)

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Ejecutar la transacción
		err = txManager.Run(context.Background(), func(ctx mongo.SessionContext) error {
			// Aquí irían las operaciones de la base de datos
			return nil
		})

		if err != nil {
			mt.Errorf("Expected no error, got %v", err)
		}
	})
}

// TestMongoTx_Run_FunctionError verifica que si la función dentro de la transacción retorna un error, este sea propagado.
func TestMongoTx_Run_FunctionError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("function error", func(mt *mtest.T) {
		// Configurar el mock para respuestas exitosas y luego abortar
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // StartTransaction
			mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    123,
				Name:    "SomeNonTransientError",
				Message: "Non-transient error occurred",
				Labels:  []string{},
			}), // CommitTransaction (fallo no transitorio)
		)

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Ejecutar la transacción que falla
		expectedErr := errors.New("function failed")
		err = txManager.Run(context.Background(), func(ctx mongo.SessionContext) error {
			return expectedErr
		})

		if err != expectedErr {
			mt.Errorf("Expected error %v, got %v", expectedErr, err)
		}
	})
}

// TestMongoTx_Run_PanicRecovery verifica que si la función provoca un pánico, este se recupera y se propaga como error.
func TestMongoTx_Run_PanicRecovery(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("panic in transaction function", func(mt *mtest.T) {
		// Configurar el mock para respuestas exitosas
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // StartTransaction
			mtest.CreateSuccessResponse(), // CommitTransaction
		)

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Ejecutar la transacción que provoca un pánico
		err = txManager.Run(context.Background(), func(ctx mongo.SessionContext) error {
			panic("unexpected panic")
		})

		expectedErrMsg := "panic occurred during transaction: unexpected panic"
		if err == nil || err.Error() != expectedErrMsg {
			mt.Errorf("Expected panic error '%s', got %v", expectedErrMsg, err)
		}
	})
}

// TestMongoTx_Run_CommitError verifica que si ocurre un error al intentar hacer commit de la transacción, este error se propaga correctamente.
func TestMongoTx_Run_CommitError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("commit transaction error", func(mt *mtest.T) {
		// Configurar el mock para simular un error en el commit de la transacción
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(),
			mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code: 8000,
				Name: "CommitTransactionError",
			}),
		)

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Ejecutar la transacción
		err = txManager.Run(context.Background(), func(ctx mongo.SessionContext) error {
			// Aquí irían las operaciones de la base de datos
			return nil
		})

		if err == nil {
			mt.Errorf("Expected commit transaction error, got nil")
		} else if err.Error() != "Failed to commit transaction" {
			mt.Errorf("Expected 'Failed to commit transaction' error, got %v", err)
		}
	})
}

// TestMongoTx_Run_AbortTransactionError verifica que si ocurre un error al intentar abortar una transacción (por ejemplo, después de un pánico), este error se maneja correctamente.
func TestMongoTx_Run_AbortTransactionError(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("abort transaction error", func(mt *mtest.T) {
		// Configurar el mock para simular un pánico y luego un error al abortar la transacción
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // StartTransaction
			mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    8002,
				Name:    "AbortTransactionError",
				Message: "Failed to abort transaction",
				Labels:  []string{},
			}), // AbortTransaction (fallo al abortar)
		)

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Ejecutar la transacción que provoca un pánico
		err = txManager.Run(context.Background(), func(ctx mongo.SessionContext) error {
			panic("unexpected panic")
		})

		expectedErrMsg := "panic occurred during transaction: unexpected panic"
		if err == nil || err.Error() != expectedErrMsg {
			mt.Errorf("Expected panic error '%s', got %v", expectedErrMsg, err)
		}
	})
}

// TestMongoTx_Run_ContextCancelledBeforeStart verifica que si el contexto es cancelado antes de iniciar la transacción, se retorna inmediatamente con un error de contexto cancelado.
func TestMongoTx_Run_ContextCancelledBeforeStart(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("context cancelled before start", func(mt *mtest.T) {
		// Crear un contexto que ya está cancelado
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Ejecutar la transacción
		err = txManager.Run(ctx, func(ctx mongo.SessionContext) error {
			// Aquí irían las operaciones de la base de datos
			return nil
		})

		if err == nil || !errors.Is(err, context.Canceled) {
			mt.Errorf("Expected context.Canceled error, got %v", err)
		}
	})
}

// TestMongoTx_Run_ContextCancelledDuringTransaction verifica que si el contexto es cancelado durante la ejecución de la transacción, se maneja correctamente abortando la transacción y retornando el error de contexto cancelado.
func TestMongoTx_Run_ContextCancelledDuringTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("context cancelled during transaction", func(mt *mtest.T) {
		// Configurar el mock para simular respuestas exitosas
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // StartTransaction
			mtest.CreateSuccessResponse(), // CommitTransaction (no se llegará aquí)
		)

		// Inicializar el TxManager
		logger := logger.New("debug")
		txManager, err := New(mt.Client, WithDefaultTxOptions(), WithLogger(logger))
		if err != nil {
			mt.Fatalf("Failed to create TxManager: %v", err)
		}

		// Crear un contexto que se cancelará durante la transacción
		ctx, cancel := context.WithCancel(context.Background())

		// Ejecutar la transacción en una goroutine
		done := make(chan error)
		go func() {
			err := txManager.Run(ctx, func(ctx mongo.SessionContext) error {
				// Simular una operación larga
				time.Sleep(2 * time.Second)
				return nil
			})
			done <- err
		}()

		// Cancelar el contexto después de 500 milisegundos
		time.Sleep(500 * time.Millisecond)
		cancel()

		// Esperar a que la transacción termine
		err = <-done

		if err == nil || !errors.Is(err, context.Canceled) {
			mt.Errorf("Expected context.Canceled error, got %v", err)
		}
	})
}
