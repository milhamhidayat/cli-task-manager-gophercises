# cli-task-manager-gophercises

CLI task manager from gophericeses task

## Installation

1. Install package

   ```sh
   go get ./..
   ```

2. Build repo
   ```sh
   go install
   ```

## List of command

| Command              | Description                                     | Example                 |
| -------------------- | ----------------------------------------------- | ----------------------- |
| task                 | Show description of task package                | `task`                  |
| task list            | Show list of task                               | `task list`             |
| task add [task]      | Add a task                                      | `task add learn golang` |
| task do [index task] | Complete a task based on **index** in task list | `task do 1`             |
| task rm [index task] | Delete a task based on **index** in task list   | `task rm 1`             |
