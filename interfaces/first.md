Task: File Viewer System

You have a system that manages files.

Requirements:
A file can:
Read content
Write content
A ViewerService should ONLY read files

Your job:
Create a struct File
Read()
Write()

Create a service:

type ViewerService struct {
    file ???
}
Ensure:
ViewerService can ONLY call Read()
Cannot call Write() (compile-time restriction)
