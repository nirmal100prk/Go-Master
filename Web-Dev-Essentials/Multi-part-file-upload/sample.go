package main

/*

	ctx := c.Request.Context()

		// TODO: file size validation to be implemented
		// decision to be made on whether file size validation is based on tenant
		filekey := app.Config.FileUploadConfig.Filekey                     // Form key for files
		totalAllowedSizeMB := app.Config.FileUploadConfig.TotalAllowedSize // Total size limit for all files in the request (in MB)
		tempBucketName := app.Config.FileUploadConfig.MinioTempBucket
		totalAllowedSizeBytes := int64(totalAllowedSizeMB * 1024 * 1024) // Convert total allowed size to bytes

		response := fieldman.PutObjectResponseDTO{
			UploadedFiles: make([]*fieldman.UploadedFileDTO, 0),
		}

		if tempBucketName == "" {
			slog.Error("tempBucketName is empty")
			response.StatusCode = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		// Parse the multipart form
		if err := c.Request.ParseMultipartForm(totalAllowedSizeBytes); err != nil {
			slog.Error("Failed to parse multipart form", slog.String("error", err.Error()))
			response.StatusCode = http.StatusBadRequest
			response.Message = "Failed to parse multipart form"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		form := c.Request.MultipartForm
		files := form.File[filekey]

		// Check if files are received
		if len(files) == 0 {
			slog.Error("No files received")
			response.StatusCode = http.StatusBadRequest
			response.Message = "No files received"
			c.JSON(http.StatusBadRequest, response)
			return
		}

		svcReq := &grpc_mussad_v1_tenant.PutObjectRequest{
			Files: make([]*grpc_mussad_v1_tenant.FileRequest, 0, len(files)),
		}

		// Iterate over each file in the request
		for _, fileHeader := range files {
			// Open the file
			file, err := fileHeader.Open()
			if err != nil {
				slog.Error("Error opening file: ", slog.String("file", fileHeader.Filename), slog.String("error", err.Error()))
				response.StatusCode = http.StatusBadRequest
				response.Message = "Failed to open file"
				c.JSON(http.StatusBadRequest, response)
				return
			}
			defer file.Close()

			// Validate file format
			contentType := fileHeader.Header.Get("Content-Type")
			if !common.IsValidFormat(contentType) {
				err := fmt.Errorf("unsupported file format: %s", contentType)
				slog.Error(err.Error())
				response.StatusCode = http.StatusUnsupportedMediaType
				response.Message = err.Error()
				c.JSON(http.StatusUnsupportedMediaType, response)
				return
			}

			// Validate file size
			// if fileHeader.Size > maxFileSizeBytes {
			// 	err := fmt.Errorf("file size exceeds the maximum limit of %d MB: %s", allowedSizeMB, fileHeader.Filename)
			// 	slog.Error(err.Error())
			// 	response.StatusCode = http.StatusBadRequest
			// 	response.Message = err.Error()
			// 	c.JSON(http.StatusBadRequest, response)
			// 	return
			// }

			// Read the file into a byte buffer
			fileBuffer := make([]byte, fileHeader.Size)
			if _, err := file.Read(fileBuffer); err != nil {
				slog.Error("Failed to read file content", slog.String("file", fileHeader.Filename), slog.String("error", err.Error()))
				response.StatusCode = http.StatusInternalServerError
				response.Message = "Failed to read file content"
				c.JSON(http.StatusInternalServerError, response)
				return
			}

			// Add file details to the service request
			svcReq.Files = append(svcReq.Files, &grpc_mussad_v1_tenant.FileRequest{
				Content:     fileBuffer,
				FileName:    fileHeader.Filename,
				ContentType: contentType,
				BucketName:  tempBucketName,
				IsMobile:    true,
			})

		}

*/
