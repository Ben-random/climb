# climb Documentation 🧗

## Overview

`climb` is a CLI tool that makes your local scripts and binaries globally available by creating aliases that can be called from anywhere in your terminal. It manages scripts by installing them to your local bin directory (`~/.local/bin` on Unix-like systems or `%LOCALAPPDATA%\Microsoft\WindowsApps` on Windows).

## Installation

See the [README](../README.md) for installation instructions.

---

## Commands

### `create`

Creates a new global alias for a script or binary.

**Syntax:**
```bash
climb create <alias> <path/to/script>
```

**Arguments:**
- `<alias>` - The name you want to use to call the script globally
- `<path/to/script>` - Path to the script or binary file

**Example:**
```bash
climb create hello ~/scripts/hello.sh
```

After running this command, you can execute `hello` from anywhere in your terminal, and it will run the script located at `~/scripts/hello.sh`.

**Behavior:**
- Validates that the provided path exists
- Checks if the alias already exists
- If alias exists, the command will fail with an error
- Copies the script to `~/.local/bin/` with the alias name
- Sets executable permissions (0755)

---

### `update`

Updates an existing alias to point to a different script or binary.

**Syntax:**
```bash
climb update <alias> <path/to/new/script>
```

**Arguments:**
- `<alias>` - The existing alias name you want to update
- `<path/to/new/script>` - Path to the new script or binary file

**Example:**
```bash
climb update hello ~/scripts/hello_v2.sh
```

**Behavior:**
- Validates that the provided path exists
- Checks if the alias exists
- If alias doesn't exist, the command will fail with an error
- Prompts for confirmation before overwriting
- Replaces the existing script with the new one

**Interactive Prompt:**
```
Do you want to override alias: hello [y/n]?
```

---

### `delete`

Removes a global alias.

**Syntax:**
```bash
climb delete <alias>
```

**Arguments:**
- `<alias>` - The name of the alias to delete

**Example:**
```bash
climb delete hello
```

**Behavior:**
- Looks up the alias in your PATH
- Displays the file location
- Prompts for confirmation before deletion
- Removes the file from your local bin directory

**Interactive Prompt:**
```
Are you sure you want to delete alias: hello? [y/n]
```

---

### `help`

Displays usage information.

**Syntax:**
```bash
climb help
```

**Output:**
```
Usage: climb <create|update|delete> <COMMAND> <path/to/script NOTE: create & update ONLY>
```

---

## Flags

### `--dry-run`

Performs a dry run without making any actual file system changes. Useful for testing commands before executing them.

**Syntax:**
```bash
climb --dry-run <command> <args>
```

**Examples:**
```bash
climb --dry-run create hello ~/scripts/hello.sh
climb --dry-run update hello ~/scripts/hello_v2.sh
climb --dry-run delete hello
```

**Behavior:**
- Prints what actions would be taken
- Does not create, modify, or delete any files
- Validates all arguments and paths
- Shows confirmation prompts (but doesn't execute the action)

**Sample Output:**
```bash
$ climb --dry-run create hello ~/scripts/hello.sh
DRY_RUN: Write file from /Users/admin/scripts/hello.sh to /Users/admin/.local/bin/hello
```

---

## Usage Examples

### Basic Workflow

1. **Create an alias for a frequently used script:**
   ```bash
   climb create deploy ~/projects/deploy-script.sh
   ```

2. **Use the alias from anywhere:**
   ```bash
   deploy
   ```

3. **Update the script to a newer version:**
   ```bash
   climb update deploy ~/projects/deploy-script-v2.sh
   ```

4. **Remove the alias when no longer needed:**
   ```bash
   climb delete deploy
   ```

### Advanced Examples

**Creating aliases for multiple scripts:**
```bash
climb create backup ~/scripts/backup.sh
climb create sync ~/scripts/sync-data.sh
climb create cleanup ~/scripts/cleanup-logs.sh
```

**Testing commands with dry-run:**
```bash
# Test before creating
climb --dry-run create myapp ~/bin/myapp

# Verify the output, then run for real
climb create myapp ~/bin/myapp
```

---

## Error Handling

### Common Errors

**"Not enough args"**
- You didn't provide required arguments
- Use `climb help` to see correct usage

**"Too many args"**
- You provided too many arguments
- Check the command syntax

**"Alias already exists"**
- Trying to create an alias that already exists
- Use `climb update` instead of `climb create`

**"Alias doesn't exist"**
- Trying to update an alias that doesn't exist
- Use `climb create` instead of `climb update`

**"Failed to find bin at path"**
- The specified script path doesn't exist
- Verify the path is correct

**"Command not found in PATH"**
- When deleting, the alias wasn't found
- The alias may have already been deleted or never existed

---

## Tips

1. **Use absolute paths** when possible to avoid confusion:
   ```bash
   climb create myalias ~/scripts/myscript.sh
   ```

2. **Test with --dry-run** before making changes:
   ```bash
   climb --dry-run create myalias ~/scripts/myscript.sh
   ```

3. **Ensure your script has a shebang** (e.g., `#!/bin/bash`) for proper execution

4. **Alias naming conventions:**
   - Use lowercase names
   - Avoid spaces
   - Keep names short and memorable
   - Avoid conflicts with existing system commands

5. **Check if an alias exists** before creating:
   ```bash
   which myalias
   ```

---

## Technical Details

### Installation Directories

- **Unix-like systems (macOS, Linux):** `~/.local/bin`
- **Windows:** `%LOCALAPPDATA%\Microsoft\WindowsApps`

### File Permissions

Scripts are installed with permissions `0755`, which means:
- Owner: read, write, execute
- Group: read, execute
- Others: read, execute

### PATH Requirements

Ensure your local bin directory is in your PATH. The installation script should handle this, but if commands aren't found, you may need to add it manually:

**For bash/zsh (~/.bashrc or ~/.zshrc):**
```bash
export PATH="$HOME/.local/bin:$PATH"
```

**For Windows:**
The WindowsApps directory is typically already in PATH.

---

## Troubleshooting

**Command not found after creation:**
1. Verify the alias was created:
   ```bash
   ls -la ~/.local/bin
   ```
2. Check if `~/.local/bin` is in your PATH:
   ```bash
   echo $PATH
   ```
3. If not in PATH, add it and restart your terminal

**Permission denied when running alias:**
1. Check file permissions:
   ```bash
   ls -la ~/.local/bin/<alias>
   ```
2. Ensure execute permission is set:
   ```bash
   chmod +x ~/.local/bin/<alias>
   ```

**Script runs but behaves incorrectly:**
1. Ensure the script has proper shebang line
2. Check if the script uses relative paths that need to be absolute
3. Verify script dependencies are available globally
