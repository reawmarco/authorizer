package graphql

import (
	"fmt"
	"strings"
)

func (g *graphqlProvider) isEmailServiceConfigured() bool {
	return strings.TrimSpace(g.Config.SMTPHost) != "" &&
		g.Config.SMTPPort > 0 &&
		strings.TrimSpace(g.Config.SMTPUsername) != "" &&
		strings.TrimSpace(g.Config.SMTPPassword) != "" &&
		strings.TrimSpace(g.Config.SMTPSenderEmail) != ""
}

func (g *graphqlProvider) sendTransactionalEmail(to []string, event string, data map[string]interface{}) error {
	log := g.Log.With().Str("func", "sendTransactionalEmail").Str("event", event).Logger()
	if !g.isEmailServiceConfigured() {
		return fmt.Errorf("email sending is disabled for this instance")
	}
	if err := g.EmailProvider.SendEmail(to, event, data); err != nil {
		log.Debug().Err(err).Msg("Failed to send transactional email")
		return fmt.Errorf("failed to send email")
	}
	return nil
}
