# **gopher-build-all**
Create all possible binaries of a project in go

## **ChangeLog**
<details>
  <summary>0.0.2</summary>
  
  ## **Flags**

  - **bin-folder-name** ─ With this flag you can rename the output directory.

  - **bin-name** ─ binary-name renamed to bin-name.

  - **only-systems** ─ only-in renamed to only-systems.

  - **sub-folder** ─ A folder is created where the binary is saved.

</details>
 
<details>
  <summary>0.0.1</summary>

  First release.

</details>

## **Install**

**Go 1.17+**

```bash
$ go install github.com/FlamesX-128/gopher-build-all@latest
```

## **Flags**
You can alter a few things when creating the binaries.

**bin-folder-name**

  - **description:** ─ The name of the folder where the binaries will be placed.

  - **default** ─ `bin`

  - **usage:** ─ `-bin-folder-name="dist"`

**bin-name**

  - **description:** ─ Name given to the binary.

  - **default** ─ `main`

  - **usage:** ─ `-bin-name="gopher-build-all"`


**max-goruntines**

  - **description:** ─ Used to indicate the maximum number of goroutines.

  - **default** ─ `3`

  - **usage:** ─ `-max-goruntines=5`

**only-systems**

  - **description:** ─ Used to indicate to only build on specific operating systems.

  - **default** ─ `""` *All possible systems*

  - **usage:** ─ `-only-systems="linux windows"`

**project-directory**

  - **description:** ─ It is used to indicate which is the project path.

  - **default** ─ `""` *The path where it was called*

  - **usage:** ─ `-project-directory="/home/User/go/github.com/User/gopher-build-all/"` *in linux*

**sub-folder**

  - **description:** ─ It is used to indicate if a sub folder should be created for the file.

  - **default** ─ `true`

  - **usage:** ─ `-sub-folder=true`
