# **nextBuild.go**
Create all possible binaries of a project in go

## **ChangeLog**
- `0.0.1` ─ First release.

## Flags
You can alter a few things when creating the binaries.

- `binary-name` ─ Rename the created file.

- `max-goruntines` ─ Maximum number of goroutines, with these the processes to create the binaries are divided.

- `only-in` ─ They only create the binaries on specific systems, like linux, windows, darwin or other.

- `project-directory` ─ This is the path where to look for the main file and it will create the folder with the binaries, by default it uses the console path.

## **Example**

### **Project**

![image](https://user-images.githubusercontent.com/78381898/143496673-c7a55734-ee4d-4cae-83ec-475eb8063293.png)

### **Command**

```bash
$ go run main.go -binary-name="nextBuild" -max-goruntines=6
```

### **Out**

![image](https://user-images.githubusercontent.com/78381898/143496842-b8f166cb-55a6-4e07-803f-ec279cbb8713.png)