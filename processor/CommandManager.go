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
	tools       []ToolCmd
}

type ToolsCommandManager struct {
	toolsCommands map[string]*ToolCmd
	groups        []toolGroup
}

func (manager *ToolsCommandManager) populateExecMap() {
	commandsMap := make(map[string]*ToolCmd)
	for _, group := range manager.groups {
		for i, _ := range group.tools {
			toolPointer := group.tools[i]
			commandsMap[toolPointer.getId()] = &toolPointer
		}
	}
	manager.toolsCommands = commandsMap
}

func (manager *ToolsCommandManager) Execute(toolsCommandId string) error {
	if value, ok := manager.toolsCommands[toolsCommandId]; ok {
		return (*value).Execute()
	} else {
		manager.UsagePrint()
		return errors.New("command not found")
	}
}

func (manager *ToolsCommandManager) UsagePrint() {
	fmt.Println(usage)
	for _, group := range manager.groups {
		fmt.Printf("* %s: ", group.description)
		fmt.Println()
		for _, tool := range group.tools {
			fmt.Println("    - ", tool.Info())
		}
		fmt.Println()
	}
}

func GetToolsCommandManager() ToolsCommandManager {
	manager := ToolsCommandManager{groups: []toolGroup{
		{
			name:        "channels_repos",
			description: "Manage channels and repositories",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-clone-by-date", "Clone channels based on a date",
					"/usr/bin/spacewalk-clone-by-date", "spacewalk-utils"),
				externalToolCommand("spacewalk-common-channels", "",
					"/usr/bin/spacewalk-common-channels", "spacewalk-utils"),
				externalToolCommand("spacewalk-manage-channel-lifecycle", "",
					"/usr/bin/spacewalk-manage-channel-lifecycle", "spacewalk-utils"),
				externalToolCommand("spacewalk-remove-channel", "",
					"/usr/bin/spacewalk-remove-channel", "spacewalk-backend-tools"),
				externalToolCommand("spacewalk-repo-sync", "",
					"/usr/bin/spacewalk-repo-sync", "spacewalk-backend-tools"),
				externalToolCommand("mgr-sync", "",
					"/usr/sbin/mgr-sync", "susemanager-tools"),
				externalToolCommand("mgr-clean-old-patchnames", "Remove patches with old patchnames from the given channels",
					"/usr/sbin/mgr-clean-old-patchnames", "susemanager-tools"),
				externalToolCommand("mgr-delete-patch", "",
					"/usr/sbin/mgr-delete-patch", "susemanager-tools"),
				// --- next 3 command are the same. symbolic link between them (/usr/bin/mgrpush -> rhnpush) (/usr/bin/rhnpush -> rhnpush-3.6)
				externalToolCommand("mgrpush", "",
					"/usr/bin/mgrpush", "mgr-push"),
				externalToolCommand("rhnpush", "",
					"/usr/bin/rhnpush", "mgr-push"),
				externalToolCommand("rhnpush-3.6", "",
					"/usr/bin/rhnpush-3.6", "python3-mgr-push"),
				externalToolCommand("mgr-sign-metadata", "",
					"/usr/bin/mgr-sign-metadata", "spacewalk-backend-tools"),
				externalToolCommand("mgr-sign-metadata-ctl", "",
					"/usr/bin/mgr-sign-metadata-ctl", "spacewalk-backend-tools"),
				externalToolCommand("spacewalk-export-channels", "",
					"/usr/bin/spacewalk-export-channels", "spacewalk-utils-extras"),
				externalToolCommand("spacewalk-watch-channel-sync.sh", "",
					"/usr/bin/spacewalk-watch-channel-sync.sh", "spacewalk-utils-extras"),
			},
		},
		{
			name:        "osBuild",
			description: "OS Image build",
			tools: []ToolCmd{
				externalToolCommand("mgr-package-rpm-certificate-osimage", "",
					"/usr/sbin/mgr-package-rpm-certificate-osimage", "spacewalk-certs-tools"),
			},
		},
		{
			name:        "serverManagement",
			description: "Server Management",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-hostname-rename", "",
					"/usr/bin/spacewalk-hostname-rename", "spacewalk-utils"),
				externalToolCommand("spacewalk-make-mount-points", "",
					"/usr/bin/spacewalk-make-mount-points", "spacewalk-setup"),
				externalToolCommand("spacewalk-service", "",
					"/usr/sbin/spacewalk-service", "spacewalk-admin"),
				externalToolCommand("spacewalk-update-signatures", "",
					"/usr/bin/spacewalk-update-signatures", "spacewalk-backend-tools"),
				externalToolCommand("rhn-sat-restart-silent", "",
					"/usr/sbin/rhn-sat-restart-silent", "spacewalk-admin"),
			},
		},
		{
			name:        "database",
			description: "Database Management",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-schema-upgrade", "",
					"/usr/bin/spacewalk-schema-upgrade", "susemanager-schema"),
				externalToolCommand("rhn-schema-stats", "",
					"/usr/bin/rhn-schema-stats", "spacewalk-backend-tools"),
				externalToolCommand("rhn-schema-version", "",
					"/usr/bin/rhn-schema-version", "spacewalk-backend-tools"),
				externalToolCommand("rhn-db-stats", "",
					"/usr/bin/rhn-db-stats", "spacewalk-backend-tools"),
				externalToolCommand("rhn-config-schema.pl", "",
					"/usr/bin/rhn-config-schema.pl", "spacewalk-admin"),
			},
		},
		{
			name:        "util",
			description: "Util",
			tools: []ToolCmd{externalToolCommand("spacecmd", "",
				"/usr/bin/spacecmd", "spacecmd"),
				externalToolCommand("spacewalk-sql", "",
					"/usr/bin/spacewalk-sql", "susemanager-schema"),
				externalToolCommand("satpasswd", "",
					"/usr/bin/satpasswd", "spacewalk-backend-tools"),
				externalToolCommand("satwho", "",
					"/usr/bin/satwho", "spacewalk-backend-tools"),
				externalToolCommand("taskotop", "",
					"/usr/bin/taskotop", "spacewalk-utils-extras"),
				externalToolCommand("rhn-search", "",
					"/usr/sbin/rhn-search", "spacewalk-search"),
				externalToolCommand("mgr-libmod", "",
					"/usr/bin/mgr-libmod", "mgr-libmod"),
				externalToolCommand("apply_errata", "",
					"/usr/bin/apply_errata", "spacewalk-utils-extras"),
				externalToolCommand("delete-old-systems-interactive", "",
					"/usr/bin/delete-old-systems-interactive", "spacewalk-utils-extras"),
				externalToolCommand("migrate-system-profile", "",
					"/usr/bin/migrate-system-profile", "spacewalk-utils-extras"),
				externalToolCommand("spacewalk-api", "",
					"/usr/bin/spacewalk-api", "spacewalk-utils-extras"),
				externalToolCommand("spacewalk-manage-snapshots", "",
					"/usr/bin/spacewalk-manage-snapshots", "spacewalk-utils-extras"),
				externalToolCommand("sw-system-snapshot", "",
					"/usr/bin/sw-system-snapshot", "spacewalk-utils-extras"),
				externalToolCommand("sw-ldap-user-sync", "",
					"/usr/bin/sw-ldap-user-sync", "spacewalk-utils-extras"),
				externalToolCommand("rhn-charsets", "",
					"/usr/bin/rhn-charsets", "spacewalk-backend-tools"),
			},
		},
		{
			name:        "serverSetup",
			description: "Server Setup",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-setup", "",
					"/usr/bin/spacewalk-setup", "spacewalk-setup"),
				externalToolCommand("spacewalk-setup-cobbler", "",
					"/usr/bin/spacewalk-setup-cobbler", "spacewalk-setup"),
				externalToolCommand("spacewalk-setup-db-ssl-certificates", "",
					"/usr/bin/spacewalk-setup-db-ssl-certificates", "spacewalk-setup"),
				externalToolCommand("spacewalk-setup-httpd", "",
					"/usr/bin/spacewalk-setup-httpd", "spacewalk-setup"),
				externalToolCommand("spacewalk-setup-ipa-authentication", "",
					"/usr/bin/spacewalk-setup-ipa-authentication", "spacewalk-setup"),
				externalToolCommand("spacewalk-setup-jabberd", "",
					"/usr/bin/spacewalk-setup-jabberd", "spacewalk-setup-jabberd"),
				externalToolCommand("spacewalk-setup-sudoers", "",
					"/usr/bin/spacewalk-setup-sudoers", "spacewalk-setup"),
				externalToolCommand("spacewalk-setup-tomcat", "",
					"/usr/bin/spacewalk-setup-tomcat", "spacewalk-setup"),
				externalToolCommand("rhn-config-satellite.pl", "",
					"/usr/bin/rhn-config-satellite.pl", "spacewalk-admin"),
			},
		},
		{
			name:        "export",
			description: "Export",
			tools: []ToolCmd{
				externalToolCommand("mgr-exporter", "",
					"/usr/bin/mgr-exporter", "spacewalk-backend-tools"),
				externalToolCommand("rhn-satellite-exporter", "",
					"/usr/bin/rhn-satellite-exporter", "spacewalk-backend-tools"),
				externalToolCommand("spacewalk-export", "",
					"/usr/bin/spacewalk-export", "spacewalk-utils-extras"),
			},
		},
		{
			name:        "cleanup",
			description: "Cleanup",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-data-fsck", "",
					"/usr/bin/spacewalk-data-fsck", "spacewalk-backend-tools"),
				externalToolCommand("spacewalk-diskcheck", "",
					"/usr/bin/spacewalk-diskcheck", "spacewalk-backend-tools"),
				externalToolCommand("spacewalk-debug", "Produce debug information",
					"/usr/bin/spacewalk-debug", "spacewalk-backend-tools"),
			},
		},
		{
			name:        "certification",
			description: "certification and report",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-fips-tool", "",
					"/usr/bin/spacewalk-fips-tool", "spacewalk-backend-tools"),
				externalToolCommand("spacewalk-report", "",
					"/usr/bin/spacewalk-report", "spacewalk-reports"),
				externalToolCommand("mgr-monitoring-ctl", "",
					"/usr/sbin/mgr-monitoring-ctl", "spacewalk-admin"),
			},
		},
		{
			name:        "Bootstrap",
			description: "Bootstrap",
			tools: []ToolCmd{
				// --- next 2 command are the same. symbolic link between them (/usr/sbin/mgr-push-register -> spacewalk-push-register)
				externalToolCommand("spacewalk-push-register", "",
					"/usr/sbin/spacewalk-push-register", "spacewalk-certs-tools"),
				externalToolCommand("mgr-push-register", "",
					"/usr/sbin/mgr-push-register", "spacewalk-certs-tools"),
				// --- next 2 command are the same. symbolic link between them (/usr/sbin/mgr-ssh-push-init -> spacewalk-ssh-push-init)
				externalToolCommand("mgr-ssh-push-init", "Setup a client to be managed via SSH push",
					"/usr/sbin/mgr-ssh-push-init", "spacewalk-certs-tools"),
				externalToolCommand("spacewalk-ssh-push-init", "Setup a client to be managed via SSH push",
					"/usr/sbin/spacewalk-ssh-push-init", "spacewalk-certs-tools"),
				// --- next 3 command are the same. symbolic link between them (/usr/bin/mgr-bootstrap -> rhn-bootstrap) (/usr/bin/rhn-bootstrap -> rhn-bootstrap-3.6)
				externalToolCommand("mgr-bootstrap", "",
					"/usr/bin/mgr-bootstrap", "spacewalk-certs-tools"),
				externalToolCommand("rhn-bootstrap", "",
					"/usr/bin/rhn-bootstrap", "spacewalk-certs-tools"),
				externalToolCommand("rhn-bootstrap-3.6", "",
					"/usr/bin/rhn-bootstrap-3.6", "python3-spacewalk-certs-tools"),
				externalToolCommand("mgr-create-bootstrap-repo", "",
					"/usr/sbin/mgr-create-bootstrap-repo", "susemanager-tools"),
			},
		},
		{
			name:        "ssl",
			description: "SSL Management",
			tools: []ToolCmd{
				// --- next 3 command are the same. symbolic link between them (//usr/bin/mgr-ssl-tool -> rhn-ssl-tool) (/usr/bin/rhn-ssl-tool -> rhn-ssl-tool-3.6)
				externalToolCommand("mgr-ssl-tool", "",
					"/usr/bin/mgr-ssl-tool", "spacewalk-certs-tools"),
				externalToolCommand("rhn-ssl-tool", "",
					"/usr/bin/rhn-ssl-tool", "spacewalk-certs-tools"),
				externalToolCommand("rhn-ssl-tool-3.6", "",
					"/usr/bin/rhn-ssl-tool-3.6", "python3-spacewalk-certs-tools"),
				// --- next 2 command are the same. symbolic link between them (/usr/bin/mgr-sudo-ssl-tool -> rhn-sudo-ssl-tool)
				externalToolCommand("mgr-sudo-ssl-tool", "",
					"/usr/bin/mgr-sudo-ssl-tool", "spacewalk-certs-tools"),
				externalToolCommand("rhn-sudo-ssl-tool", "",
					"/usr/bin/rhn-sudo-ssl-tool", "spacewalk-certs-tools"),
				externalToolCommand("rhn-install-ssl-cert.pl", "",
					"/usr/bin/rhn-install-ssl-cert.pl", "spacewalk-admin"),
				externalToolCommand("rhn-ssl-dbstore", "",
					"/usr/bin/rhn-ssl-dbstore", "spacewalk-backend-tools"),
				externalToolCommand("rhn-deploy-ca-cert.pl", "",
					"/usr/bin/rhn-deploy-ca-cert.pl", "spacewalk-admin"),
				externalToolCommand("rhn-generate-pem.pl", "",
					"/usr/bin/rhn-generate-pem.pl", "spacewalk-admin"),
			},
		},
		{
			name:        "Proxy",
			description: "Proxy",
			tools: []ToolCmd{
				//	// --- next 2 command are the same. symbolic link between them (/usr/sbin/rhn-profile-sync -> rhn-profile-sync-3.6)
				externalToolCommand("rhn-profile-sync", "",
					"/usr/sbin/rhn-profile-sync", "spacewalk-client-tools"),
				externalToolCommand("rhn-profile-sync-3.6", "",
					"/usr/sbin/rhn-profile-sync-3.6", "python3-spacewalk-client-tools"),
			},
		},
		{
			name:        "iss",
			description: "Inter Server Sync",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-sync-setup", "",
					"/usr/bin/spacewalk-sync-setup", "spacewalk-utils"),
				// --- next 2 command are the same. symbolic link between them (/usr/bin/mgr-inter-sync -> satellite-sync)
				externalToolCommand("satellite-sync", "",
					"/usr/bin/satellite-sync", "spacewalk-backend-tools"),
				externalToolCommand("mgr-inter-sync", "",
					"/usr/bin/mgr-inter-sync", "spacewalk-backend-tools"),
			},
		},
		{
			name:        "others",
			description: "others",
			tools: []ToolCmd{
				externalToolCommand("spacewalk-cfg-get", "",
					"/usr/bin/spacewalk-cfg-get", "spacewalk-backend"),
				externalToolCommand("spacewalk-startup-helper", "",
					"/usr/sbin/spacewalk-startup-helper",
					"spacewalk-admin"),
				externalToolCommand("satcon-build-dictionary.pl", "",
					"/usr/bin/satcon-build-dictionary.pl",
					"perl-Satcon-4.1.1-411.2.22.devel41.noarch"),
				externalToolCommand("satcon-deploy-tree.pl", "",
					"/usr/bin/satcon-deploy-tree.pl", "perl-Satcon"),
				externalToolCommand("rhn-satellite", "",
					"/usr/sbin/rhn-satellite", "spacewalk-admin"),
				externalToolCommand("rhn-satellite-activate", "",
					"/usr/bin/rhn-satellite-activate", "spacewalk-backend-tools"),
			},
		},
	}}

	manager.populateExecMap()
	return manager
}
