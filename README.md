# mycli

## Command Override Behavior

By default, if a command exists in both `mycli` and `commoncli`, the version from **commoncli** will be used.

If you want your own (`mycli`) commands to override those from `commoncli` (when there is a name conflict), set the environment variable `COMMAND_OVERRIDE` to `mycli` before running the CLI.

### Usage

**Default (commoncli wins on conflicts):**
```powershell
.\mycli.exe
```

**To let mycli commands override commoncli on conflicts:**
```powershell
$env:COMMAND_OVERRIDE="mycli"
.\mycli.exe
```

Or in Command Prompt (cmd.exe):
```cmd
set COMMAND_OVERRIDE=mycli
mycli.exe
```

If you want this override to always apply, set the environment variable permanently in your system/user