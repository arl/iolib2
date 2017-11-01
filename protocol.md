Iolib2 protocol description
###########################

# Commands


## SET-PORT

`SET-PORT|configuration_string`

### description
Set Output Port with "configuration_string". The port should not already be
configured when RENEW-PORT is received of an error will be returned.

## RESET

`RESET|`

### description
does nothing if no port is configured
close the current port if the port is configured


## WRITE

`WRITE|string`

### description
Add the string `string` to the transmit buffer of iolib.


## SEND

`SEND|`

# description
Send the content of iolib transmit buffer


# Return Values

Every command send to iolib2 should be only aknowledged after reception of the
return values

string of the form: "int|string"

where:
 - int is the error code (0 success)
 - string is the error description

If the port crashes:
 - int = -1
 - string is the stack trace
