package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Types
// FileInfo holds information about a file or directory
type FileInfo struct {
	Name         string
	Size         int64
	Mode         os.FileMode
	ModTime      time.Time
	IsDir        bool
	Permissions  string
	Owner        string
	Group        string
	LinkTarget   string
	AbsolutePath string
}
type FileType string

// GetFileHashes returns the MD5, SHA1, and SHA256 hashes of a file
func GetFileHashes(path string) (string, string, string, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return "", "", "", err
	}
	defer file.Close()

	// Create hashers for MD5, SHA1, and SHA256
	md5Hasher := md5.New()
	sha1Hasher := sha1.New()
	sha256Hasher := sha256.New()

	// Copy the file contents to the hashers
	if _, err := io.Copy(md5Hasher, file); err != nil {
		return "", "", "", err
	}
	file.Seek(0, 0) // Reset the file pointer to the beginning
	if _, err := io.Copy(sha1Hasher, file); err != nil {
		return "", "", "", err
	}
	file.Seek(0, 0) // Reset the file pointer to the beginning
	if _, err := io.Copy(sha256Hasher, file); err != nil {
		return "", "", "", err
	}

	// Get the hashes as hex strings
	md5Hash := fmt.Sprintf("%x", md5Hasher.Sum(nil))
	sha1Hash := fmt.Sprintf("%x", sha1Hasher.Sum(nil))
	sha256Hash := fmt.Sprintf("%x", sha256Hasher.Sum(nil))

	return md5Hash, sha1Hash, sha256Hash, nil
}

// GetFileType determines the file type based on the file extension
func GetFileType(path string) FileType {
	ext := filepath.Ext(path)
	// Define a map of file extensions to file type decriptions
	fileTypeMap := map[string]FileType{
		".txt":    "Text File",
		".jpg":    "Image File",
		".png":    "Image File",
		".mp4":    "Video File",
		".mp3":    "Audio File",
		".pdf":    "PDF Document",
		".docx":   "Word Document",
		".xlsx":   "Excel Spreadsheet",
		".pptx":   "PowerPoint Presentation",
		".eml":    "Email File",
		".msg":    "Email File",
		".csv":    "CSV File",
		".html":   "HTML File",
		".css":    "CSS File",
		".js":     "JavaScript File",
		".py":     "Python File",
		".rtf":    "RTF Document",
		".exe":    "Windows Executable File",
		".apk":    "Android Package File",
		".zip":    "ZIP Archive",
		".tar":    "TAR Archive",
		".gz":     "GZIP Archive",
		".bz2":    "BZIP2 Archive",
		".7z":     "7-Zip Archive",
		".rar":    "RAR Archive",
		".iso":    "ISO Image File",
		".dmg":    "Disk Image File",
		".dll":    "Dynamic Link Library",
		".so":     "Shared Object File",
		".class":  "Java Class File",
		".jar":    "Java Archive File",
		".go":     "Go Source File",
		".sh":     "Shell Script",
		".bat":    "Batch File",
		".ps1":    "PowerShell Script",
		".json":   "JSON File",
		".xml":    "XML File",
		".yaml":   "YAML File",
		".yml":    "YAML File",
		".svg":    "SVG File",
		".woff":   "Web Open Font Format",
		".woff2":  "Web Open Font Format 2",
		".ttf":    "TrueType Font",
		".otf":    "OpenType Font",
		".eot":    "Embedded OpenType Font",
		".fnt":    "Bitmap Font",
		".ttc":    "TrueType Collection",
		".pdb":    "Program Database File",
		".mdb":    "Microsoft Access Database",
		".sqlite": "SQLite Database",
		".db":     "Database File",
		".log":    "Log File",
		".tmp":    "Temporary File",
		".bak":    "Backup File",
		".swp":    "Swap File",
		".swo":    "Swap File",
		".lock":   "Lock File",
		".pid":    "Process ID File",
		".seed":   "Seed File",
		".key":    "Key File",
		".pem":    "Privacy Enhanced Mail File",
		".crt":    "Certificate File",
		".cer":    "Certificate File",
		".csr":    "Certificate Signing Request",
		".p12":    "PKCS#12 File",
		".pfx":    "PKCS#12 File",
		".p7b":    "PKCS#7 Certificate File",
		".p7c":    "PKCS#7 Certificate File",
		".p7s":    "PKCS#7 Signature File",
		".p8":     "PKCS#8 Private Key File",
		".jks":    "Java KeyStore File",
		".jceks":  "Java Cryptography Extension KeyStore File",
	}
	if fileType, exists := fileTypeMap[ext]; exists {
		return fileType
	}
	return "Unknown File Type"
}

// GetFileInfo returns information about a file or directory
func GetFileInfo(path string) (FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return FileInfo{}, err
	}
	owner, group, err := GetOwnerAndGroup(path)
	if err != nil {
		return FileInfo{}, err
	}
	permissions := fileInfo.Mode().Perm()
	fileType := GetFileType(path)
	linkTarget := ""
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		linkTarget, err = os.Readlink(path)
		if err != nil {
			return FileInfo{}, err
		}
	}
	return FileInfo{
		Name:         fileInfo.Name(),
		Size:         fileInfo.Size(),
		Mode:         fileInfo.Mode(),
		ModTime:      fileInfo.ModTime(),
		IsDir:        fileInfo.IsDir(),
		Permissions:  permissions.String(),
		Owner:        owner,
		Group:        group,
		LinkTarget:   linkTarget,
		AbsolutePath: absPath,
	}, nil
}
