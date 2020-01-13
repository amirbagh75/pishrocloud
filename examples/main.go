package main

import "github.com/amirbagh75/pishrocloud"

func main() {

	p := pishrocloud.Storage{
		APIKey:   "XXXXXXXXXXXXXXXXXXXXXX",
		AuthURL:  "http://213.233.176.12:5000/v3/auth/tokens/", // based on this docs: https://blog.pishrocloud.com/pishro-object-storage-doc/
		SwiftURL: "http://213.233.176.12:8080/swift/v1/",       // based on this docs: https://blog.pishrocloud.com/pishro-object-storage-doc/
		UserName: "YYYYYYYYYYYY",                               // your pishrocloud panel username: https://pishrocloud.com/authentication/login
		PassWord: "ZZZZZZZZZZZZ",                               // your pishrocloud panel password: https://pishrocloud.com/authentication/login
	}

	// refresh token
	// by default, token exist for 24 hours and you should refresh it every 24h on your program
	if p.RefreshToken() == true {
		// create container
		containerName := "foo"
		if p.CreateContainer(containerName) == true {
			println("container " + containerName + " created successfully.")
		} else {
			println("container " + containerName + " exist!")
		}

		// upload object to container
		localFilePath := "/tmp/upload.mp4"
		objectName := "uuu.mp4"

		// add optional metadata
		headers := map[string]string{
			"X-Object-Meta-username": "amribagh75",
			"X-Object-Meta-name":     "amir",
			"X-Object-Meta-id":       "123456789",
		}
		if (p.UploadObject(localFilePath, objectName, containerName, headers)) == true {
			println("file from \"" + localFilePath + "\" uploaded successfully with name: " + objectName)
		} else {
			println("file can't uploaded")
		}

		// donwload object to container
		localFilePath = "/tmp/dwonload.mp4"
		objectName = "uuu.mp4"
		if (p.DownloadObject(localFilePath, objectName, containerName)) == true {
			println("file \"" + objectName + "\" downloaded successfully and saved in: " + localFilePath)
		} else {
			println("file can't downloaded")
		}

		// check object exist or not
		println(p.IsObjectExist(objectName, containerName))
	} else {
		println("can't refresh token!")
	}

}
