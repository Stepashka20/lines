# Lines
This is a simple command-line tool that counts the number of lines of code in a project. It uses the cobra library to provide a user-friendly interface and allows you to specify which files to ignore while counting.

## 🖥 Usage
To use the line counter, you will need to have Go installed on your system. Once you have Go installed, you can run the following command to download the package:


```
git clone https://github.com/Stepashka20/lines
go install
```
Next, you can navigate to the root of any of your projects in the terminal and run the following command:


```
lines
```
By default, the tool will start counting lines from the current directory and all its subdirectories. You can also specify specific directories that do not need to be scanned.

```
lines --exclude=file1,file2,file3
```

## 🚩 Flags

| Short command | Long command | Description |
| --- | --- | --- |
| --info | -i | Display information about each file while scanning. This is additional information that shows the scanned file and the number of lines in it |
| --exclude | -e | Specify files to exclude from the line count. Files should be separated by a comma |

## ℹ Example:
```
lines
```
![image](https://user-images.githubusercontent.com/40739871/211648922-5ae50372-d9ea-4ae3-9651-1166636237ab.png)

```
lines -i
```
![image](https://user-images.githubusercontent.com/40739871/211648997-be25c487-1bd3-44bf-ad85-42b893aa8c61.png)
```
lines -i -e themes.zip
```
![image](https://user-images.githubusercontent.com/40739871/211649073-a20c7283-0b00-4429-871c-77b123bf92fc.png)


## 👥 Contribute
If you find any bugs or have any suggestions for improvements, please feel free to open an issue or a pull request on the GitHub repository.

## 📄 License
This project is licensed under the MIT License.
