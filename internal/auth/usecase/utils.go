package usecase

import (
	"os"

	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/internal/auth/entity"
	"gitlab.com/JorgeO3/flowcast/internal/auth/service"
)

func createUserEntity(input UserRegistrationInput) (*entity.User, error) {
	return entity.NewUser(
		input.Username,
		input.Email,
		input.Password,
		entity.WithFullName(input.FullName),
		entity.WithBirthdate(input.Birthdate),
		entity.WithGender(input.Gender),
		entity.WithPhone(input.Phone),
	)
}

func createUserPrefEntity(input UserRegistrationInput, userID int) *entity.UserPref {
	return entity.NewUserPref(userID, input.EmailNotif, input.SMSNotif)
}

func createEmailVerificationToken(userID int) (*entity.EmailVerificationToken, error) {
	return entity.NewEmailVerificationToken(userID)
}

func createMailerConfig(cfg *configs.AuthConfig, user *entity.User, emailVer *entity.EmailVerificationToken) (*service.MailerConfig, error) {
	byteHTMLTemplate, err := os.ReadFile(cfg.EmailTemplate)
	if err != nil {
		return nil, err
	}

	htmlTemplate := string(byteHTMLTemplate)

	data := map[string]string{
		"UserName":        user.FullName,
		"ConfirmationURL": emailVer.Token,
	}

	return service.NewMailerConfig(data, user.Email, htmlTemplate, "email_template"), nil
}

func createSocialLinkEntities(input UserRegistrationInput, userID int) []*entity.SocialLink {
	socialLinksLen := len(input.SocialLinks)
	socialLinks := make([]*entity.SocialLink, socialLinksLen)

	for i, socialLink := range input.SocialLinks {
		socialLinks[i] = entity.NewSocialLink(userID, socialLink.Platform, socialLink.URL)
	}
	return socialLinks
}
