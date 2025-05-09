## Sec-Utils
A go-based CLI application providing fundamental utilities to the cybersecurity community using GNU/Unix design principles.

## Contact 
clemiondev@gmail.com

## Usage
~~~
sec-utils [options] [utility]
~~~

## Features/Utilities
- info [Utility]: gather information about file or directory relevant to security practicioners and display info on the command line, or optionally output to json file as well.
    - [file] file size, file hashes like SHA256, file contents type (PE, ELF, PDF...), VirusTotal report url
    - [directory] full directory path, number of files, list of contained files, contained file sha256 hashes
    Usage (blank or . in the file/dir place equals current working directory): 
    ~~~
    sec-utils [options] info [file or dir] 
    ~~~
    Options: --output=output.json (saves json data to a specified name in the current working directory, or a subdirectory if specified like --output=subdir/output.json)
- iocs [Utility]: extract indicators of compromise from a text-based file and display them using json format on the command line, or output that data to a json file as well. (example file-types: .eml, .msg, .txt, .pdf, .html) (iocs does not work with binary formats like an executable, this is planned to change in future releases)
    - [file]    