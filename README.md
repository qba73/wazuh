# wazuh

`wazuh` is a Go client library for [Wazuh](https://wazuh.com) - The Open Source Security Platform API.

`wazuh` allows interaction with the Wazuh manager from any Go application, including web services.

`wazctl` is a command-line tool utilizing the `wazuh` library. It allows you to interact with the Wazuch manager from a command line or any script on the Linux, Windows, or MacOS platforms.

The goal of the Go library and the CLI tool is to accommodate remote management of the Wazuch manager via the command line to efficiently perform everyday tasks such as adding an agent, restarting the manager(s) or agent(s), or looking up system check details.
