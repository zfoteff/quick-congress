package cli

import (
	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/provider"
)

type CLIInterface struct {
	provider.CongressProvider
	provider.BillProvider
}

/*
 * Setup the CLI interface provider to connect to the congress API client
 */
func (c *CLIInterface) initProvider() {
	client := client.NewCongressClient()
	c.CongressProvider = provider.NewCongressProvider(client)
}

/*
* Get current Congress session from Congress endpoint
 */
func (c *CLIInterface) GetCurrentCongressSession() (current_session string) {
	c.initProvider()
	response := c.GetCurrentCongress()
	return response.ToString()
}

/*
* Get congress session by session number
 */
func (c *CLIInterface) GetCongressSession(session uint16) (selected_session string) {
	c.initProvider()
	response := c.GetCongress(session)
	return response.Congress.ToString()
}

/*
* Get Congress sessions by limit and offset from Congress endpoint.
 */
func (c *CLIInterface) GetCongressSessions(limit uint16, offset uint16) (sessions string) {
	c.initProvider()
	response := c.GetCongresses(limit, offset)
	return response.ToString()
}
