# OTP VERIFICATION
- One-Time Password process is a widely used mechanism for verifying a user's identity  during online transactions.
- Hashing OTP stored in database for verifying OTP is a standard practice for protection against database breach, unauthorised access.
- Most OTPs are 4 to 8 digits long, with numeric codes being the most common. Some systems may use alphanumeric OTPs for added complexity.
- To prevent abuse, some systems introduce a slight delay between OTP resend requests, encouraging users to wait and reducing the load on SMS gateways or email servers.


	##STEPS
	- User Request for OTP 
		- requesting otp with mobile number, email
		- The backend  will check the configuration settings to validate whether verification is required and which method is allowed
	- OTP Generation
		- using cryptographic random number generator (alphanumeric OTPs)
		- short lifespan
		- single use
		- hash and store in DB
	- OTP Delivery
		- delivery confirmation
	- User Entry of OTP
	- OTP Validation
		- Server-side Verification
			- check validity of time window
			- Track number of incorrect attempts
			- the user entered input hashed and compared with stored otp hash
			- invalidate otp after successful match verification
			
	- Handling Resends and Retries
		- Resend mechanism 
			- if the user doesnot receive the otp or it expires before they can use it, they should have option to request new OTP
			- Should impose limits on how many times OTP can be requested within a specific time frame
		- Rate limiting
			- enforce rate limits on OTP requests
		- Notification on maximum attempts