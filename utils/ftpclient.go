package ftpclient

import (
	"fmt"
	"io"
	"os"
	"time"

	"file-collector/config"

	"github.com/jlaffaye/ftp"
)

func DownloadFile(conn *ftp.ServerConn, remotePath, localPath string) error {
	// Open the local file for writing
	localFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close() // Ensure closure regardless of error

	// Retrieve the remote file (returns a ReadCloser stream)
	var resp *ftp.Response // Declare resp before using it
	resp, err = conn.Retr(remotePath)
	if err != nil {
		return err
	}
	defer resp.Close() // Ensure stream is closed at the end

	// Copy data from the remote file to the local file
	_, err = io.Copy(localFile, resp)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded file: %s\n", localPath)
	return nil
}

func ListFiles(conn *ftp.ServerConn, remoteDir string) ([]string, error) {
	entries, err := conn.List(remoteDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFile {
			files = append(files, entry.Name)
		}
	}

	return files, nil
}

func EstablishConnection(conf config.Config) (*ftp.ServerConn, error) {
	// Create FTP connection
	conn, err := ftp.Dial(fmt.Sprintf("%s:21", conf.Host), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	err = conn.Login(conf.Username, conf.Password)
	if err != nil {
		conn.Quit()
		return nil, err
	}

	return conn, nil
}
