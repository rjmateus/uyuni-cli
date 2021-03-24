package processor

import (
	"errors"
	"fmt"
)

const (
	usage = `uyuni tools 
Usage should be "uyuni-cli [command]"

Available Command:`
)

type toolGroup struct {
	name        string
	description string
	tools       []*ToolCmd
}

type ToolsCommandManager struct {
	toolsCommands map[string]ToolCmd
	groupsOrdered []string
	groups        map[string]toolGroup
}

func (manager *ToolsCommandManager) Execute(toolsCommandId string) error {
	if value, ok := manager.toolsCommands[toolsCommandId]; ok {
		return value.Execute()
	} else {
		manager.UsagePrint()
		return errors.New("command not found")
	}
}

func (manager *ToolsCommandManager) UsagePrint() {
	fmt.Println(usage)
	for _, groupID := range manager.groupsOrdered {
		group := manager.groups[groupID]
		fmt.Printf("* %s: ", group.description)
		fmt.Println()
		for _, tool := range group.tools {
			fmt.Println("    - ", (*tool).Info())
		}
		fmt.Println()
	}
}

func (manager *ToolsCommandManager) addGroup(groupId string, groupDesc string) toolGroup {
	manager.groupsOrdered = append(manager.groupsOrdered, groupId)
	manager.groups[groupId] = toolGroup{groupId, groupDesc, make([]*ToolCmd, 0)}
	return manager.groups[groupId]
}

func (manager *ToolsCommandManager) registerToolInGroup(groupId string, tool ToolCmd) {
	manager.toolsCommands[tool.getId()] = tool
	group, ok := manager.groups[groupId]
	if !ok {
		group = manager.addGroup(groupId, groupId)
	}
	group.tools = append(group.tools, &tool)
	manager.groups[groupId] = group
}

func populateChannelsrepos(manager *ToolsCommandManager) {
	groupId := "channels_repos"
	groupDesc := "Manage channels and repositories"
	manager.addGroup(groupId, groupDesc)

	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-clone-by-date", "/usr/bin/spacewalk-clone-by-date", "Clone channels based on a date", "spacewalk-utils-4.1.13-411.5.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-common-channels", "/usr/bin/spacewalk-common-channels", "", "spacewalk-utils-4.1.13-411.5.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-manage-channel-lifecycle", "/usr/bin/spacewalk-manage-channel-lifecycle", "", "spacewalk-utils-4.1.13-411.5.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-remove-channel", "/usr/bin/spacewalk-remove-channel", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-repo-sync", "/usr/bin/spacewalk-repo-sync", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-sync", "/usr/sbin/mgr-sync", "", "susemanager-tools-4.1.23-411.1.19.devel41.x86_64"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-clean-old-patchnames", "/usr/sbin/mgr-clean-old-patchnames", "Remove patches with old patchnames from the given channels", "susemanager-tools-4.1.23-411.1.19.devel41.x86_64"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-delete-patch", "/usr/sbin/mgr-delete-patch", "", "susemanager-tools-4.1.23-411.1.19.devel41.x86_64"))
	// --- next 3 command are the same. symbolic link between them (/usr/bin/mgrpush -> rhnpush) (/usr/bin/rhnpush -> rhnpush-3.6)
	manager.registerToolInGroup(groupId, externalToolCommand("mgrpush", "/usr/bin/mgrpush", "", "mgr-push-4.1.1-411.2.85.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhnpush", "/usr/bin/rhnpush", "", "mgr-push-4.1.1-411.2.85.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhnpush-3.6", "/usr/bin/rhnpush-3.6", "", "python3-mgr-push-4.1.1-411.2.85.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-sign-metadata", "/usr/bin/mgr-sign-metadata", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-sign-metadata-ctl", "/usr/bin/mgr-sign-metadata-ctl", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-export-channels", "/usr/bin/spacewalk-export-channels", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-watch-channel-sync.sh", "/usr/bin/spacewalk-watch-channel-sync.sh", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
}

func populateOsBuild(manager *ToolsCommandManager) {
	groupId := "osBuild"
	groupDesc := "OS Image build"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-package-rpm-certificate-osimage", "/usr/sbin/mgr-package-rpm-certificate-osimage", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
}

func populateServerManagement(manager *ToolsCommandManager) {
	groupId := "serverManagement"
	groupDesc := "Server Management"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-hostname-rename", "/usr/bin/spacewalk-hostname-rename", "", "spacewalk-utils-4.1.13-411.5.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-make-mount-points", "/usr/bin/spacewalk-make-mount-points", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-service", "/usr/sbin/spacewalk-service", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-update-signatures", "/usr/bin/spacewalk-update-signatures", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-sat-restart-silent", "/usr/sbin/rhn-sat-restart-silent", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
}

func populateDatabaseSchema(manager *ToolsCommandManager) {
	groupId := "database"
	groupDesc := "Database Management"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-schema-upgrade", "/usr/bin/spacewalk-schema-upgrade", "", "susemanager-schema-4.1.18-411.5.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-schema-stats", "/usr/bin/rhn-schema-stats", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-schema-version", "/usr/bin/rhn-schema-version", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-db-stats", "/usr/bin/rhn-db-stats", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-config-schema.pl", "/usr/bin/rhn-config-schema.pl", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
}

func populateUtil(manager *ToolsCommandManager) {
	groupId := "util"
	groupDesc := "Util"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-sql", "/usr/bin/spacewalk-sql", "", "susemanager-schema-4.1.18-411.5.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("satpasswd", "/usr/bin/satpasswd", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("satwho", "/usr/bin/satwho", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("taskotop", "/usr/bin/taskotop", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-search", "/usr/sbin/rhn-search", "", "spacewalk-search-4.1.4-411.1.27.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-libmod", "/usr/bin/mgr-libmod", "", "mgr-libmod-4.1.6-411.1.2.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("apply_errata", "/usr/bin/apply_errata", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("delete-old-systems-interactive", "/usr/bin/delete-old-systems-interactive", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("migrate-system-profile", "/usr/bin/migrate-system-profile", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-api", "/usr/bin/spacewalk-api", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-manage-snapshots", "/usr/bin/spacewalk-manage-snapshots", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("sw-system-snapshot", "/usr/bin/sw-system-snapshot", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("sw-ldap-user-sync", "/usr/bin/sw-ldap-user-sync", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-charsets", "/usr/bin/rhn-charsets", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
}

func populateServerSetup(manager *ToolsCommandManager) {
	groupId := "serverSetup"
	groupDesc := "Server Setup"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup", "/usr/bin/spacewalk-setup", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-cobbler", "/usr/bin/spacewalk-setup-cobbler", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-db-ssl-certificates", "/usr/bin/spacewalk-setup-db-ssl-certificates", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-httpd", "/usr/bin/spacewalk-setup-httpd", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-ipa-authentication", "/usr/bin/spacewalk-setup-ipa-authentication", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-jabberd", "/usr/bin/spacewalk-setup-jabberd", "", "spacewalk-setup-jabberd-4.1.1-411.2.26.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-sudoers", "/usr/bin/spacewalk-setup-sudoers", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-setup-tomcat", "/usr/bin/spacewalk-setup-tomcat", "", "spacewalk-setup-4.1.7-411.1.8.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-config-satellite.pl", "/usr/bin/rhn-config-satellite.pl", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
}

func populateExport(manager *ToolsCommandManager) {
	groupId := "export"
	groupDesc := "Export"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-exporter", "/usr/bin/mgr-exporter", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-satellite-exporter", "/usr/bin/rhn-satellite-exporter", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-export", "/usr/bin/spacewalk-export", "", "spacewalk-utils-extras-4.1.13-411.7.6.devel41.noarch"))
}

func populateCleanup(manager *ToolsCommandManager) {
	groupId := "cleanup"
	groupDesc := "Cleanup"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-data-fsck", "/usr/bin/spacewalk-data-fsck", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-diskcheck", "/usr/bin/spacewalk-diskcheck", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-debug", "/usr/bin/spacewalk-debug", "Produce debug information", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
}

func populateCertification(manager *ToolsCommandManager) {
	groupId := "certification"
	groupDesc := "certification and report"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-fips-tool", "/usr/bin/spacewalk-fips-tool", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-report", "/usr/bin/spacewalk-report", "", "spacewalk-reports-4.1.3-411.1.3.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-monitoring-ctl", "/usr/sbin/mgr-monitoring-ctl", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
}

func populateBootstrap(manager *ToolsCommandManager) {
	groupId := "Bootstrap"
	groupDesc := "Bootstrap"
	manager.addGroup(groupId, groupDesc)
	// --- next 2 command are the same. symbolic link between them (/usr/sbin/mgr-push-register -> spacewalk-push-register)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-push-register", "/usr/sbin/spacewalk-push-register", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-push-register", "/usr/sbin/mgr-push-register", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	// --- next 2 command are the same. symbolic link between them (/usr/sbin/mgr-ssh-push-init -> spacewalk-ssh-push-init)
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-ssh-push-init", "/usr/sbin/mgr-ssh-push-init", "Setup a client to be managed via SSH push", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-ssh-push-init", "/usr/sbin/spacewalk-ssh-push-init", "Setup a client to be managed via SSH push", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	// --- next 3 command are the same. symbolic link between them (/usr/bin/mgr-bootstrap -> rhn-bootstrap) (/usr/bin/rhn-bootstrap -> rhn-bootstrap-3.6)
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-bootstrap", "/usr/bin/mgr-bootstrap", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-bootstrap", "/usr/bin/rhn-bootstrap", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-bootstrap-3.6", "/usr/bin/rhn-bootstrap-3.6", "", "python3-spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-create-bootstrap-repo", "/usr/sbin/mgr-create-bootstrap-repo", "", "susemanager-tools-4.1.23-411.1.19.devel41.x86_64"))
}

func populateSSL(manager *ToolsCommandManager) {
	groupId := "ssl"
	groupDesc := "SSL Management"
	manager.addGroup(groupId, groupDesc)
	// --- next 3 command are the same. symbolic link between them (//usr/bin/mgr-ssl-tool -> rhn-ssl-tool) (/usr/bin/rhn-ssl-tool -> rhn-ssl-tool-3.6)
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-ssl-tool", "/usr/bin/mgr-ssl-tool", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-ssl-tool", "/usr/bin/rhn-ssl-tool", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-ssl-tool-3.6", "/usr/bin/rhn-ssl-tool-3.6", "", "python3-spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	// --- next 2 command are the same. symbolic link between them (/usr/bin/mgr-sudo-ssl-tool -> rhn-sudo-ssl-tool)
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-sudo-ssl-tool", "/usr/bin/mgr-sudo-ssl-tool", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-sudo-ssl-tool", "/usr/bin/rhn-sudo-ssl-tool", "", "spacewalk-certs-tools-4.1.14-411.1.16.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-install-ssl-cert.pl", "/usr/bin/rhn-install-ssl-cert.pl", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-ssl-dbstore", "/usr/bin/rhn-ssl-dbstore", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-deploy-ca-cert.pl", "/usr/bin/rhn-deploy-ca-cert.pl", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-generate-pem.pl", "/usr/bin/rhn-generate-pem.pl", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
}

func populateProxy(manager *ToolsCommandManager) {
	groupId := "Proxy"
	groupDesc := "Proxy"
	manager.addGroup(groupId, groupDesc)
	// --- next 2 command are the same. symbolic link between them (/usr/sbin/rhn-profile-sync -> rhn-profile-sync-3.6)
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-profile-sync", "/usr/sbin/rhn-profile-sync", "", "spacewalk-client-tools-4.1.8-411.1.17.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-profile-sync-3.6", "/usr/sbin/rhn-profile-sync-3.6", "", "python3-spacewalk-client-tools-4.1.8-411.1.17.devel41.noarch"))
}

func populateISS(manager *ToolsCommandManager) {
	groupId := "iss"
	groupDesc := "Inter Server Sync"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-sync-setup", "/usr/bin/spacewalk-sync-setup", "", "spacewalk-utils-4.1.13-411.5.1.devel41.noarch"))
	// --- next 2 command are the same. symbolic link between them (/usr/bin/mgr-inter-sync -> satellite-sync)
	manager.registerToolInGroup(groupId, externalToolCommand("satellite-sync", "/usr/bin/satellite-sync", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("mgr-inter-sync", "/usr/bin/mgr-inter-sync", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
}

func populateOthers(manager *ToolsCommandManager) {
	groupId := "others"
	groupDesc := "Others"
	manager.addGroup(groupId, groupDesc)
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-cfg-get", "/usr/bin/spacewalk-cfg-get", "", "spacewalk-backend-4.1.20-411.6.1.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewalk-startup-helper", "/usr/sbin/spacewalk-startup-helper", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("satcon-build-dictionary.pl", "/usr/bin/satcon-build-dictionary.pl", "", "perl-Satcon-4.1.1-411.2.22.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("satcon-deploy-tree.pl", "/usr/bin/satcon-deploy-tree.pl", "", "perl-Satcon-4.1.1-411.2.22.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-satellite", "/usr/sbin/rhn-satellite", "", "spacewalk-admin-4.1.8-411.1.19.devel41.noarch"))
	manager.registerToolInGroup(groupId, externalToolCommand("spacewealk-final-archive", "", "", ""))
	manager.registerToolInGroup(groupId, externalToolCommand("rhn-satellite-activate", "/usr/bin/rhn-satellite-activate", "", "spacewalk-backend-tools-4.1.20-411.6.1.devel41.noarch"))
}

func GetToolsCommandManager() ToolsCommandManager {
	manager := ToolsCommandManager{make(map[string]ToolCmd), make([]string, 0), make(map[string]toolGroup)}
	populateChannelsrepos(&manager)
	populateOsBuild(&manager)
	populateServerManagement(&manager)
	populateDatabaseSchema(&manager)
	populateUtil(&manager)
	populateServerSetup(&manager)
	populateExport(&manager)
	populateCleanup(&manager)
	populateCertification(&manager)
	populateBootstrap(&manager)
	populateSSL(&manager)
	populateProxy(&manager)
	populateISS(&manager)
	populateOthers(&manager)
	return manager
}
