# Getting Started
## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
	* **Go**  (Visit https://golang.org/doc/install for more details on how to install Go)

	* **Java VM** Version 8 or later

	* **Eclipse 4.6 (Neon)** or later (http://www.eclipse.org/neon/ )

	* **GoClipse** setup within above installed Eclipse (Follow the instructions at https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions to successfully setup GoClipse)

* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.

* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:
   ```Go
   go get github.com/apimatic/unirest-go
   ```
	This will install unirest-go in the ```GOPATH``` you specified in the system variables.

* Now follow the below mentioned steps to build your SDK:
	1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

		![Building SDK - Step 1](http://apidocs.io/illustration/go?step=build0&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

	2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

		![Building SDK - Step 2](http://apidocs.io/illustration/go?step=build1&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

	3. In ```Select root directory```, provide path to the unzipped archive for the generated code

		![Building SDK - Step 3](http://apidocs.io/illustration/go?step=build2&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

		Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

		![Building SDK - Step 4](http://apidocs.io/illustration/go?step=build3&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

	4.	The Go library will be imported and its files will be visible in the Project Explorer

		![Building SDK - Step 5](http://apidocs.io/illustration/go?step=build4&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

## How to Use

The following section explains how to use the TestMessage360Simpl2 library in a new project.

#### 1. Add a new Test Project
Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse - Step 1](http://apidocs.io/illustration/go?step=createNewProject0&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Name the Project as ```Test``` and click ```Finish```

![Add a new project in Eclipse - Step 2](http://apidocs.io/illustration/go?step=createNewProject1&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Create a new directory in the ```src``` directory of this project

![Add a new project in Eclipse - Step 3](http://apidocs.io/illustration/go?step=createNewProject2&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Name it ```test.com```

![Add a new project in Eclipse - Step 4](http://apidocs.io/illustration/go?step=createNewProject3&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Now create a new file inside ```src/test.com```

![Add a new project in Eclipse - Step 5](http://apidocs.io/illustration/go?step=createNewProject4&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Name it ```testsdk.go```

![Add a new project in Eclipse - Step 6](http://apidocs.io/illustration/go?step=createNewProject5&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.


#### 2. Configure the Test Project
You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Configure the Test Project - Step 1](http://apidocs.io/illustration/go?step=configureProject0&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Configure the Test Project - Step 2](http://apidocs.io/illustration/go?step=configureProject1&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Append the library path to this GOPATH

![Configure the Test Project - Step 3](http://apidocs.io/illustration/go?step=configureProject2&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

Once the path is appended, click on ```OK```

![Configure the Test Project - Step 4](http://apidocs.io/illustration/go?step=configureProject3&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)


#### 3. Build the Test Project
Right click on the project name and click on ```Build Project```

![Build Test Project - Step 1](http://apidocs.io/illustration/go?step=buildProject0&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)


#### 4. Run the Test Project
If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Test Project - Step 1](http://apidocs.io/illustration/go?step=runProject0&workspaceFolder=TestMessage360_simpl2-GoLang&workspaceName=TestMessage360Simpl2&projectName=testmessage360simpl2_lib)

## Initialization

#### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| accountSid | Account sid for authentication |
| authToken | Auth token for authentication  |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference
## <a name="list_of_controllers"></a>List of Controllers

* [carrier_pkg](#carrier_pkg)
* [call_pkg](#call_pkg)
* [sms_pkg](#sms_pkg)
* [account_pkg](#account_pkg)
* [conference_pkg](#conference_pkg)
* [phonenumber_pkg](#phonenumber_pkg)
* [recording_pkg](#recording_pkg)
* [transcription_pkg](#transcription_pkg)
* [usage_pkg](#usage_pkg)

## <a name="carrier_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".carrier_pkg") carrier_pkg


#### Get instance
Factory for the ``` CARRIER ``` interface can be accessed from the package carrier_pkg.
```go
carrier := carrier_pkg.NewCARRIER()
```

### <a name="create_carrier_lookup_xml"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookupXML") CreateCarrierLookupXML

> Get the Carrier lookup in XML 

```go
func (me *CARRIER_IMPL) CreateCarrierLookupXML(phonenumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | The Number to lookup |



#### Example Usage:
```go
phonenumber := "phonenumber"

var result string
result,_ = carrier.CreateCarrierLookupXML(phonenumber)

```





### <a name="create_carrier_lookup_json"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookupJSON") CreateCarrierLookupJSON

> Get the Carrier Lookup Json

```go
func (me *CARRIER_IMPL) CreateCarrierLookupJSON(phonenumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | The number to lookup |



#### Example Usage:
```go
phonenumber := "phonenumber"

var result string
result,_ = carrier.CreateCarrierLookupJSON(phonenumber)

```





### <a name="create_carrier_lookup_list_xml"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookupListXML") CreateCarrierLookupListXML

> Get the All Purchase Number's Carrier lookup in XML

```go
func (me *CARRIER_IMPL) CreateCarrierLookupListXML(
            page *string,
            pagesize *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page Number |
| pagesize |  ``` Optional ```  | Page Size |



#### Example Usage:
```go
page := "page"
pagesize := "pagesize"

var result string
result,_ = carrier.CreateCarrierLookupListXML(page, pagesize)

```





### <a name="create_carrier_lookup_list_json"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookupListJSON") CreateCarrierLookupListJSON

> Get the All Purchase Number's Carrier lookup in JSON

```go
func (me *CARRIER_IMPL) CreateCarrierLookupListJSON(
            page *string,
            pagesize *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page Number |
| pagesize |  ``` Optional ```  | Page Size |



#### Example Usage:
```go
page := "page"
pagesize := "pagesize"

var result string
result,_ = carrier.CreateCarrierLookupListJSON(page, pagesize)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="call_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".call_pkg") call_pkg


#### Get instance
Factory for the ``` CALL ``` interface can be accessed from the package call_pkg.
```go
call := call_pkg.NewCALL()
```

### <a name="create_view_call_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateViewCallXML") CreateViewCallXML

> View Call Response XML

```go
func (me *CALL_IMPL) CreateViewCallXML(callsid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | Call Sid id for particular Call |



#### Example Usage:
```go
callsid := "callsid"

var result string
result,_ = call.CreateViewCallXML(callsid)

```





### <a name="create_view_call_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateViewCallJSON") CreateViewCallJSON

> View Call Response JSON

```go
func (me *CALL_IMPL) CreateViewCallJSON(callsid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | Call Sid id for particular Call |



#### Example Usage:
```go
callsid := "callsid"

var result string
result,_ = call.CreateViewCallJSON(callsid)

```





### <a name="create_record_call_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateRecordCallXML") CreateRecordCallXML

> Record Call XML

```go
func (me *CALL_IMPL) CreateRecordCallXML(
            callSid string,
            record bool,
            direction models_pkg.DirectionEnum,
            timeLimit *int64,
            callBackUrl *string,
            fileformat models_pkg.AudioFormatEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| record |  ``` Required ```  | Set true to initiate recording or false to terminate recording |
| direction |  ``` Optional ```  | The leg of the call to record |
| timeLimit |  ``` Optional ```  | Time in seconds the recording duration should not exceed |
| callBackUrl |  ``` Optional ```  | URL consulted after the recording completes |
| fileformat |  ``` Optional ```  | Format of the recording file. Can be .mp3 or .wav |



#### Example Usage:
```go
callSid := "CallSid"
record := true
direction := models_pkg.Direction_IN
timeLimit,_ := strconv.ParseInt("136", 10, 8)
callBackUrl := "CallBackUrl"
fileformat := models_pkg.AudioFormat_MP3

var result string
result,_ = call.CreateRecordCallXML(callSid, record, direction, timeLimit, callBackUrl, fileformat)

```





### <a name="create_record_call_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateRecordCallJSON") CreateRecordCallJSON

> Record Call Json

```go
func (me *CALL_IMPL) CreateRecordCallJSON(
            callSid string,
            record bool,
            direction models_pkg.DirectionEnum,
            timeLimit *int64,
            callBackUrl *string,
            fileformat models_pkg.AudioFormatEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| record |  ``` Required ```  | Set true to initiate recording or false to terminate recording |
| direction |  ``` Optional ```  | The leg of the call to record |
| timeLimit |  ``` Optional ```  | Time in seconds the recording duration should not exceed |
| callBackUrl |  ``` Optional ```  | URL consulted after the recording completes |
| fileformat |  ``` Optional ```  | Format of the recording file. Can be .mp3 or .wav |



#### Example Usage:
```go
callSid := "CallSid"
record := true
direction := models_pkg.Direction_IN
timeLimit,_ := strconv.ParseInt("136", 10, 8)
callBackUrl := "CallBackUrl"
fileformat := models_pkg.AudioFormat_MP3

var result string
result,_ = call.CreateRecordCallJSON(callSid, record, direction, timeLimit, callBackUrl, fileformat)

```





### <a name="create_voice_effect_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateVoiceEffectXML") CreateVoiceEffectXML

> Voice Effect JSON

```go
func (me *CALL_IMPL) CreateVoiceEffectXML(
            callSid string,
            audioDirection models_pkg.AudioDirectionEnum,
            pitchSemiTones *float64,
            pitchOctaves *float64,
            pitch *float64,
            rate *float64,
            tempo *float64)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | TODO: Add a parameter description |
| audioDirection |  ``` Optional ```  | TODO: Add a parameter description |
| pitchSemiTones |  ``` Optional ```  | value between -14 and 14 |
| pitchOctaves |  ``` Optional ```  | value between -1 and 1 |
| pitch |  ``` Optional ```  | value greater than 0 |
| rate |  ``` Optional ```  | value greater than 0 |
| tempo |  ``` Optional ```  | value greater than 0 |



#### Example Usage:
```go
callSid := "CallSid"
audioDirection := models_pkg.AudioDirection_IN
pitchSemiTones := 136.1149825743
pitchOctaves := 136.1149825743
pitch := 136.1149825743
rate := 136.1149825743
tempo := 136.1149825743

var result string
result,_ = call.CreateVoiceEffectXML(callSid, audioDirection, pitchSemiTones, pitchOctaves, pitch, rate, tempo)

```





### <a name="create_voice_effect_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateVoiceEffectJSON") CreateVoiceEffectJSON

> Voice Effect JSON

```go
func (me *CALL_IMPL) CreateVoiceEffectJSON(
            callSid string,
            audioDirection models_pkg.AudioDirectionEnum,
            pitchSemiTones *float64,
            pitchOctaves *float64,
            pitch *float64,
            rate *float64,
            tempo *float64)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | TODO: Add a parameter description |
| audioDirection |  ``` Optional ```  | TODO: Add a parameter description |
| pitchSemiTones |  ``` Optional ```  | value between -14 and 14 |
| pitchOctaves |  ``` Optional ```  | value between -1 and 1 |
| pitch |  ``` Optional ```  | value greater than 0 |
| rate |  ``` Optional ```  | value greater than 0 |
| tempo |  ``` Optional ```  | value greater than 0 |



#### Example Usage:
```go
callSid := "CallSid"
audioDirection := models_pkg.AudioDirection_IN
pitchSemiTones := 136.1149825743
pitchOctaves := 136.1149825743
pitch := 136.1149825743
rate := 136.1149825743
tempo := 136.1149825743

var result string
result,_ = call.CreateVoiceEffectJSON(callSid, audioDirection, pitchSemiTones, pitchOctaves, pitch, rate, tempo)

```





### <a name="create_play_audio_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreatePlayAudioXML") CreatePlayAudioXML

> Play audio XML response

```go
func (me *CALL_IMPL) CreatePlayAudioXML(
            callSid string,
            audioUrl string,
            length *int64,
            direction models_pkg.DirectionEnum,
            loop *bool,
            mix *bool)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| audioUrl |  ``` Required ```  | URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav |
| length |  ``` Optional ```  | Time limit in seconds for audio play back |
| direction |  ``` Optional ```  | The leg of the call audio will be played to |
| loop |  ``` Optional ```  | Repeat audio playback indefinitely |
| mix |  ``` Optional ```  | If false, all other audio will be muted |



#### Example Usage:
```go
callSid := "CallSid"
audioUrl := "AudioUrl"
length,_ := strconv.ParseInt("136", 10, 8)
direction := models_pkg.Direction_IN
loop := true
mix := true

var result string
result,_ = call.CreatePlayAudioXML(callSid, audioUrl, length, direction, loop, mix)

```





### <a name="create_play_audio_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreatePlayAudioJSON") CreatePlayAudioJSON

> Play Dtmf and send the Digit IN JSON

```go
func (me *CALL_IMPL) CreatePlayAudioJSON(
            length int64,
            direction models_pkg.DirectionEnum,
            loop bool,
            mix bool,
            callSid *string,
            audioUrl *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| length |  ``` Required ```  | Time limit in seconds for audio play back |
| direction |  ``` Required ```  | The leg of the call audio will be played to |
| loop |  ``` Required ```  | Repeat audio playback indefinitely |
| mix |  ``` Required ```  | If false, all other audio will be muted |
| callSid |  ``` Optional ```  | The unique identifier of each call resource |
| audioUrl |  ``` Optional ```  | URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav |



#### Example Usage:
```go
length,_ := strconv.ParseInt("136", 10, 8)
direction := models_pkg.Direction_IN
loop := true
mix := true
callSid := "CallSid"
audioUrl := "AudioUrl"

var result string
result,_ = call.CreatePlayAudioJSON(length, direction, loop, mix, callSid, audioUrl)

```





### <a name="create_list_call_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateListCallXML") CreateListCallXML

> Returned xml response here contains a list of calls associated with your Message360 account

```go
func (me *CALL_IMPL) CreateListCallXML(
            page *string,
            pageSize *string,
            to *string,
            from *string,
            dateCreated *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Only list calls to this number |
| from |  ``` Optional ```  | Only list calls from this number |
| dateCreated |  ``` Optional ```  | Only list calls starting within the specified date range |



#### Example Usage:
```go
page := "Page"
pageSize := "PageSize"
to := "To"
from := "From"
dateCreated := "DateCreated"

var result string
result,_ = call.CreateListCallXML(page, pageSize, to, from, dateCreated)

```





### <a name="create_list_call_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateListCallJSON") CreateListCallJSON

> Returned Json response here contains a list of calls associated with your Message360 account

```go
func (me *CALL_IMPL) CreateListCallJSON(
            page *string,
            pageSize *string,
            to *string,
            from *string,
            dateCreated *string)(,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Only list calls to this number |
| from |  ``` Optional ```  | Only list calls from this number |
| dateCreated |  ``` Optional ```  | Only list calls starting within the specified date range |



#### Example Usage:
```go
page := "Page"
pageSize := "PageSize"
to := "To"
from := "From"
dateCreated := "DateCreated"

var result 
result,_ = call.CreateListCallJSON(page, pageSize, to, from, dateCreated)

```





### <a name="create_interrupted_call_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateInterruptedCallJSON") CreateInterruptedCallJSON

> Interrupt the Call by Call Sid

```go
func (me *CALL_IMPL) CreateInterruptedCallJSON(
            callSid string,
            url *string,
            method models_pkg.HttpMethodEnum,
            status models_pkg.InterruptedCallStatusEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | Call SId |
| url |  ``` Optional ```  | URL the in-progress call will be redirected to |
| method |  ``` Optional ```  | The method used to request the above Url parameter |
| status |  ``` Optional ```  | Status to set the in-progress call to |



#### Example Usage:
```go
callSid := "CallSid"
url := "Url"
method := models_pkg.HttpMethod_GET
status := models_pkg.InterruptedCallStatus_CANCELED

var result string
result,_ = call.CreateInterruptedCallJSON(callSid, url, method, status)

```





### <a name="create_send_digit_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateSendDigitJSON") CreateSendDigitJSON

> Play Dtmf and send the Digit IN JSON

```go
func (me *CALL_IMPL) CreateSendDigitJSON(
            callSid string,
            playDtmf string,
            playDtmfDirection models_pkg.DirectionEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |



#### Example Usage:
```go
callSid := "CallSid"
playDtmf := "PlayDtmf"
playDtmfDirection := models_pkg.Direction_IN

var result string
result,_ = call.CreateSendDigitJSON(callSid, playDtmf, playDtmfDirection)

```





### <a name="create_interrupted_call_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateInterruptedCallXML") CreateInterruptedCallXML

> Interrupt the Call by Call Sid XML Format

```go
func (me *CALL_IMPL) CreateInterruptedCallXML(
            callSid string,
            url *string,
            method models_pkg.HttpMethodEnum,
            status models_pkg.InterruptedCallStatusEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | Call Sid |
| url |  ``` Optional ```  | URL the in-progress call will be redirected to |
| method |  ``` Optional ```  | The method used to request the above Url parameter |
| status |  ``` Optional ```  | Status to set the in-progress call to |



#### Example Usage:
```go
callSid := "CallSid"
url := "Url"
method := models_pkg.HttpMethod_GET
status := models_pkg.InterruptedCallStatus_CANCELED

var result string
result,_ = call.CreateInterruptedCallXML(callSid, url, method, status)

```





### <a name="create_send_digit_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateSendDigitXML") CreateSendDigitXML

> Play Dtmf and send the Digit in XML 

```go
func (me *CALL_IMPL) CreateSendDigitXML(
            callSid string,
            playDtmf string,
            playDtmfDirection models_pkg.DirectionEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |



#### Example Usage:
```go
callSid := "CallSid"
playDtmf := "PlayDtmf"
playDtmfDirection := models_pkg.Direction_IN

var result string
result,_ = call.CreateSendDigitXML(callSid, playDtmf, playDtmfDirection)

```





### <a name="create_make_call_xml"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateMakeCallXML") CreateMakeCallXML

> You can experiment with initiating a call through Message360 and view the request response generated when doing so and get the response in xml

```go
func (me *CALL_IMPL) CreateMakeCallXML(
            method models_pkg.HttpMethodEnum,
            statusCallBackUrl string,
            statusCallBackMethod models_pkg.HttpMethodEnum,
            fallBackUrl string,
            fallBackMethod models_pkg.HttpMethodEnum,
            heartBeatUrl string,
            heartBeatMethod models_pkg.HttpMethodEnum,
            timeout int64,
            playDtmf string,
            hideCallerId bool,
            record bool,
            recordCallBackUrl string,
            recordCallBackMethod models_pkg.HttpMethodEnum,
            transcribe bool,
            transcribeCallBackUrl string,
            ifMachine models_pkg.IfMachineEnum,
            fromCountryCode *string,
            from *string,
            toCountryCode *string,
            to *string,
            url *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| method |  ``` Required ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Required ```  | specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| statusCallBackMethod |  ``` Required ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Required ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Required ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| heartBeatUrl |  ``` Required ```  | URL that can be requested every 60 seconds during the call to notify of elapsed time |
| heartBeatMethod |  ``` Required ```  | Specifies the HTTP method used to request HeartbeatUrl. |
| timeout |  ``` Required ```  | Time (in seconds) Message360 should wait while the call is ringing before canceling the call. |
| playDtmf |  ``` Required ```  | DTMF Digits to play to the call once it connects. 0-9, #, or * |
| hideCallerId |  ``` Required ```  | Specifies if the caller id will be hidden |
| record |  ``` Required ```  | Specifies if the call should be recorded |
| recordCallBackUrl |  ``` Required ```  | Recording parameters will be sent here upon completion |
| recordCallBackMethod |  ``` Required ```  | Method used to request the RecordCallback URL. |
| transcribe |  ``` Required ```  | Specifies if the call recording should be transcribed |
| transcribeCallBackUrl |  ``` Required ```  | Transcription parameters will be sent here upon completion |
| ifMachine |  ``` Required ```  | How Message360 should handle the receiving numbers voicemail machine |
| fromCountryCode |  ``` Optional ```  | from country code |
| from |  ``` Optional ```  | TODO: Add a parameter description |
| toCountryCode |  ``` Optional ```  | To cuntry code number |
| to |  ``` Optional ```  | To number |
| url |  ``` Optional ```  | URL requested once the call connects |



#### Example Usage:
```go
method := models_pkg.HttpMethod_GET
statusCallBackUrl := "StatusCallBackUrl"
statusCallBackMethod := models_pkg.HttpMethod_GET
fallBackUrl := "FallBackUrl"
fallBackMethod := models_pkg.HttpMethod_GET
heartBeatUrl := "HeartBeatUrl"
heartBeatMethod := models_pkg.HttpMethod_GET
timeout,_ := strconv.ParseInt("136", 10, 8)
playDtmf := "PlayDtmf"
hideCallerId := true
record := true
recordCallBackUrl := "RecordCallBackUrl"
recordCallBackMethod := models_pkg.HttpMethod_GET
transcribe := true
transcribeCallBackUrl := "TranscribeCallBackUrl"
ifMachine := models_pkg.ifMachine_CONTINUE
fromCountryCode := "FromCountryCode"
from := "From"
toCountryCode := "ToCountryCode"
to := "To"
url := "Url"

var result string
result,_ = call.CreateMakeCallXML(method, statusCallBackUrl, statusCallBackMethod, fallBackUrl, fallBackMethod, heartBeatUrl, heartBeatMethod, timeout, playDtmf, hideCallerId, record, recordCallBackUrl, recordCallBackMethod, transcribe, transcribeCallBackUrl, ifMachine, fromCountryCode, from, toCountryCode, to, url)

```





### <a name="create_make_call_json"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateMakeCallJSON") CreateMakeCallJSON

> You can experiment with initiating a call through Message360 and view the request response generated when doing so and get the response in json

```go
func (me *CALL_IMPL) CreateMakeCallJSON(
            fromCountryCode string,
            from string,
            toCountryCode string,
            to string,
            url string,
            method models_pkg.HttpMethodEnum,
            statusCallBackUrl *string,
            statusCallBackMethod models_pkg.HttpMethodEnum,
            fallBackUrl *string,
            fallBackMethod models_pkg.HttpMethodEnum,
            heartBeatUrl *string,
            heartBeatMethod *bool,
            timeout *int64,
            playDtmf *string,
            hideCallerId *bool,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod models_pkg.HttpMethodEnum,
            transcribe *bool,
            transcribeCallBackUrl *string,
            ifMachine models_pkg.IfMachineEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromCountryCode |  ``` Required ```  | from country code |
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| toCountryCode |  ``` Required ```  | To cuntry code number |
| to |  ``` Required ```  | To number |
| url |  ``` Required ```  | URL requested once the call connects |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| heartBeatUrl |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed tim |
| heartBeatMethod |  ``` Optional ```  | Specifies the HTTP method used to request HeartbeatUrl. |
| timeout |  ``` Optional ```  | Time (in seconds) Message360 should wait while the call is ringing before canceling the call |
| playDtmf |  ``` Optional ```  | DTMF Digits to play to the call once it connects. 0-9, #, or * |
| hideCallerId |  ``` Optional ```  | Specifies if the caller id will be hidden |
| record |  ``` Optional ```  | Specifies if the call should be recorded |
| recordCallBackUrl |  ``` Optional ```  | Recording parameters will be sent here upon completion |
| recordCallBackMethod |  ``` Optional ```  | Method used to request the RecordCallback URL. |
| transcribe |  ``` Optional ```  | Specifies if the call recording should be transcribed |
| transcribeCallBackUrl |  ``` Optional ```  | Transcription parameters will be sent here upon completion |
| ifMachine |  ``` Optional ```  | How Message360 should handle the receiving numbers voicemail machine |



#### Example Usage:
```go
fromCountryCode := "FromCountryCode"
from := "From"
toCountryCode := "ToCountryCode"
to := "To"
url := "Url"
method := models_pkg.HttpMethod_GET
statusCallBackUrl := "StatusCallBackUrl"
statusCallBackMethod := models_pkg.HttpMethod_GET
fallBackUrl := "FallBackUrl"
fallBackMethod := models_pkg.HttpMethod_GET
heartBeatUrl := "HeartBeatUrl"
heartBeatMethod := true
timeout,_ := strconv.ParseInt("136", 10, 8)
playDtmf := "PlayDtmf"
hideCallerId := true
record := true
recordCallBackUrl := "RecordCallBackUrl"
recordCallBackMethod := models_pkg.HttpMethod_GET
transcribe := true
transcribeCallBackUrl := "TranscribeCallBackUrl"
ifMachine := models_pkg.ifMachine_CONTINUE

var result string
result,_ = call.CreateMakeCallJSON(fromCountryCode, from, toCountryCode, to, url, method, statusCallBackUrl, statusCallBackMethod, fallBackUrl, fallBackMethod, heartBeatUrl, heartBeatMethod, timeout, playDtmf, hideCallerId, record, recordCallBackUrl, recordCallBackMethod, transcribe, transcribeCallBackUrl, ifMachine)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="sms_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".sms_pkg") sms_pkg


#### Get instance
Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.
```go
sMS := sms_pkg.NewSMS()
```

### <a name="create_view_smsxml"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateViewSMSXML") CreateViewSMSXML

> View Particular SMS  XML Response

```go
func (me *SMS_IMPL) CreateViewSMSXML(messagesid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| messagesid |  ``` Required ```  | Message Sid |



#### Example Usage:
```go
messagesid := "messagesid"

var result string
result,_ = sMS.CreateViewSMSXML(messagesid)

```





### <a name="create_view_smsjson"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateViewSMSJSON") CreateViewSMSJSON

> View Particular SMS  JSON Response

```go
func (me *SMS_IMPL) CreateViewSMSJSON(messagesid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| messagesid |  ``` Required ```  | Message sid |



#### Example Usage:
```go
messagesid := "messagesid"

var result string
result,_ = sMS.CreateViewSMSJSON(messagesid)

```





### <a name="create_list_in_bound_smsxml"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListInBoundSMSXML") CreateListInBoundSMSXML

> List All Inbound SMS in XML format

```go
func (me *SMS_IMPL) CreateListInBoundSMSXML(
            page *int64,
            pagesize *int64,
            from *string,
            to *string)(,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound SMS |
| to |  ``` Optional ```  | To Number to get Inbound SMS |



#### Example Usage:
```go
page,_ := strconv.ParseInt("136", 10, 8)
pagesize,_ := strconv.ParseInt("136", 10, 8)
from := "from"
to := "to"

var result 
result,_ = sMS.CreateListInBoundSMSXML(page, pagesize, from, to)

```





### <a name="create_list_in_bound_smsjson"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListInBoundSMSJSON") CreateListInBoundSMSJSON

> List All Inbound SMS in JSON format

```go
func (me *SMS_IMPL) CreateListInBoundSMSJSON(
            page *int64,
            pagesize *string,
            from *string,
            to *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound SMS |
| to |  ``` Optional ```  | To Number to get Inbound SMS |



#### Example Usage:
```go
page,_ := strconv.ParseInt("136", 10, 8)
pagesize := "pagesize"
from := "from"
to := "to"

var result string
result,_ = sMS.CreateListInBoundSMSJSON(page, pagesize, from, to)

```





### <a name="create_list_smsxml"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListSMSXML") CreateListSMSXML

> List All SMS in XML format

```go
func (me *SMS_IMPL) CreateListSMSXML(
            page *int64,
            pagesize *string,
            from *string,
            to *string,
            datesent *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Messages sent from this number |
| to |  ``` Optional ```  | Messages sent to this number |
| datesent |  ``` Optional ```  | Messages sent in the specified date range |



#### Example Usage:
```go
page,_ := strconv.ParseInt("136", 10, 8)
pagesize := "pagesize"
from := "from"
to := "to"
datesent := "datesent"

var result string
result,_ = sMS.CreateListSMSXML(page, pagesize, from, to, datesent)

```





### <a name="create_list_smsjson"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListSMSJSON") CreateListSMSJSON

> List All SMS in JSON format

```go
func (me *SMS_IMPL) CreateListSMSJSON(
            page *int64,
            pagesize *int64,
            from *string,
            to *string,
            datesent *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Messages sent from this number |
| to |  ``` Optional ```  | Messages sent to this number |
| datesent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |



#### Example Usage:
```go
page,_ := strconv.ParseInt("136", 10, 8)
pagesize,_ := strconv.ParseInt("136", 10, 8)
from := "from"
to := "to"
datesent := "datesent"

var result string
result,_ = sMS.CreateListSMSJSON(page, pagesize, from, to, datesent)

```





### <a name="create_send_smsxml"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateSendSMSXML") CreateSendSMSXML

> Send SMS through Message360 XML

```go
func (me *SMS_IMPL) CreateSendSMSXML(
            fromcountrycode int64,
            from string,
            tocountrycode int64,
            to string,
            body string,
            method models_pkg.HttpMethodEnum,
            messagestatuscallback *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromcountrycode |  ``` Required ```  ``` DefaultValue ```  | From Country Code |
| from |  ``` Required ```  | SMS enabled Message360 number to send this message from |
| tocountrycode |  ``` Required ```  ``` DefaultValue ```  | To country code |
| to |  ``` Required ```  | Number to send the SMS to |
| body |  ``` Required ```  | Text Message To Send |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messagestatuscallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |



#### Example Usage:
```go
fromcountrycode,_ := strconv.ParseInt("1", 10, 8)
from := "from"
tocountrycode,_ := strconv.ParseInt("1", 10, 8)
to := "to"
body := "body"
method := models_pkg.HttpMethod_GET
messagestatuscallback := "messagestatuscallback"

var result string
result,_ = sMS.CreateSendSMSXML(fromcountrycode, from, tocountrycode, to, body, method, messagestatuscallback)

```





### <a name="create_send_smsjson"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateSendSMSJSON") CreateSendSMSJSON

> Send SMS through Message360 Json

```go
func (me *SMS_IMPL) CreateSendSMSJSON(
            fromcountrycode int64,
            from string,
            tocountrycode int64,
            to string,
            body string,
            method models_pkg.HttpMethodEnum,
            messagestatuscallback *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromcountrycode |  ``` Required ```  ``` DefaultValue ```  | From Country Code |
| from |  ``` Required ```  | SMS enabled Message360 number to send this message from |
| tocountrycode |  ``` Required ```  ``` DefaultValue ```  | To country code |
| to |  ``` Required ```  | Number to send the SMS to |
| body |  ``` Required ```  | Text Message To Send |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messagestatuscallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |



#### Example Usage:
```go
fromcountrycode,_ := strconv.ParseInt("1", 10, 8)
from := "from"
tocountrycode,_ := strconv.ParseInt("1", 10, 8)
to := "to"
body := "body"
method := models_pkg.HttpMethod_GET
messagestatuscallback := "messagestatuscallback"

var result string
result,_ = sMS.CreateSendSMSJSON(fromcountrycode, from, tocountrycode, to, body, method, messagestatuscallback)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="account_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".account_pkg") account_pkg


#### Get instance
Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.
```go
account := account_pkg.NewACCOUNT()
```

### <a name="create_view_account_x_ml"></a>![Method: ](http://apidocs.io/img/method.png ".account_pkg.CreateViewAccountXMl") CreateViewAccountXMl

> View Account response in xml format

```go
func (me *ACCOUNT_IMPL) CreateViewAccountXMl(date string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | TODO: Add a parameter description |



#### Example Usage:
```go
date := "date"

var result string
result,_ = account.CreateViewAccountXMl(date)

```





### <a name="create_view_account_json"></a>![Method: ](http://apidocs.io/img/method.png ".account_pkg.CreateViewAccountJson") CreateViewAccountJson

> Display Account Description

```go
func (me *ACCOUNT_IMPL) CreateViewAccountJson(date string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | TODO: Add a parameter description |



#### Example Usage:
```go
date := "date"

var result string
result,_ = account.CreateViewAccountJson(date)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="conference_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".conference_pkg") conference_pkg


#### Get instance
Factory for the ``` CONFERENCE ``` interface can be accessed from the package conference_pkg.
```go
conference := conference_pkg.NewCONFERENCE()
```

### <a name="create_view_conference_json"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewConferenceJSON") CreateViewConferenceJSON

> View Conference JSON

```go
func (me *CONFERENCE_IMPL) CreateViewConferenceJSON(conferencesid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | The unique identifier of each conference resource |



#### Example Usage:
```go
conferencesid := "conferencesid"

var result string
result,_ = conference.CreateViewConferenceJSON(conferencesid)

```





### <a name="create_view_conference_xml"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewConferenceXML") CreateViewConferenceXML

> View Conference XML

```go
func (me *CONFERENCE_IMPL) CreateViewConferenceXML(conferencesid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | The unique identifier of each conference resource |



#### Example Usage:
```go
conferencesid := "conferencesid"

var result string
result,_ = conference.CreateViewConferenceXML(conferencesid)

```





### <a name="create_view_participant_json"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewParticipantJSON") CreateViewParticipantJSON

> View Participant JSON

```go
func (me *CONFERENCE_IMPL) CreateViewParticipantJSON(
            conferenceSid string,
            participantSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| participantSid |  ``` Required ```  | TODO: Add a parameter description |



#### Example Usage:
```go
conferenceSid := "ConferenceSid"
participantSid := "ParticipantSid"

var result string
result,_ = conference.CreateViewParticipantJSON(conferenceSid, participantSid)

```





### <a name="create_view_participant_xml"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewParticipantXML") CreateViewParticipantXML

> View Participant XML

```go
func (me *CONFERENCE_IMPL) CreateViewParticipantXML(
            conferenceSid string,
            participantSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| participantSid |  ``` Required ```  | TODO: Add a parameter description |



#### Example Usage:
```go
conferenceSid := "ConferenceSid"
participantSid := "ParticipantSid"

var result string
result,_ = conference.CreateViewParticipantXML(conferenceSid, participantSid)

```





### <a name="create_list_conference_json"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListConferenceJSON") CreateListConferenceJSON

> List Conference JSON

```go
func (me *CONFERENCE_IMPL) CreateListConferenceJSON(
            page *int64,
            pageSize *int64,
            friendlyName *string,
            status models_pkg.InterruptedCallStatusEnum,
            dateCreated *string,
            dateUpdated *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| dateUpdated |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("136", 10, 8)
pageSize,_ := strconv.ParseInt("136", 10, 8)
friendlyName := "FriendlyName"
status := models_pkg.InterruptedCallStatus_CANCELED
dateCreated := "DateCreated"
dateUpdated := "DateUpdated"

var result string
result,_ = conference.CreateListConferenceJSON(page, pageSize, friendlyName, status, dateCreated, dateUpdated)

```





### <a name="create_list_conference_xml"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListConferenceXML") CreateListConferenceXML

> List Conference JSON

```go
func (me *CONFERENCE_IMPL) CreateListConferenceXML(
            page *int64,
            pageSize *int64,
            friendlyName *string,
            status models_pkg.InterruptedCallStatusEnum,
            dateCreated *string,
            dateUpdated *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| dateUpdated |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("136", 10, 8)
pageSize,_ := strconv.ParseInt("136", 10, 8)
friendlyName := "FriendlyName"
status := models_pkg.InterruptedCallStatus_CANCELED
dateCreated := "DateCreated"
dateUpdated := "DateUpdated"

var result string
result,_ = conference.CreateListConferenceXML(page, pageSize, friendlyName, status, dateCreated, dateUpdated)

```





### <a name="add_participant_json"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.AddParticipantJSON") AddParticipantJSON

> Add Participant in conference JSON 

```go
func (me *CONFERENCE_IMPL) AddParticipantJSON(
            conferencesid string,
            participantnumber string,
            tocountrycode int64,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | Unique Conference Sid |
| participantnumber |  ``` Required ```  | Particiant Number |
| tocountrycode |  ``` Required ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
conferencesid := "conferencesid"
participantnumber := "participantnumber"
tocountrycode,_ := strconv.ParseInt("136", 10, 8)
muted := true
deaf := true

var result string
result,_ = conference.AddParticipantJSON(conferencesid, participantnumber, tocountrycode, muted, deaf)

```





### <a name="add_participant_xml"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.AddParticipantXML") AddParticipantXML

> Add Participant in conference XML 

```go
func (me *CONFERENCE_IMPL) AddParticipantXML(
            conferencesid string,
            participantnumber string,
            tocountrycode int64,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | Unique Conference Sid |
| participantnumber |  ``` Required ```  | Particiant Number |
| tocountrycode |  ``` Required ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
conferencesid := "conferencesid"
participantnumber := "participantnumber"
tocountrycode,_ := strconv.ParseInt("136", 10, 8)
muted := true
deaf := true

var result string
result,_ = conference.AddParticipantXML(conferencesid, participantnumber, tocountrycode, muted, deaf)

```





### <a name="create_list_participant_json"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListParticipantJSON") CreateListParticipantJSON

> List Participant JSON

```go
func (me *CONFERENCE_IMPL) CreateListParticipantJSON(
            conferenceSid string,
            page *int64,
            pagesize *int64,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| page |  ``` Optional ```  | page number |
| pagesize |  ``` Optional ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
conferenceSid := "ConferenceSid"
page,_ := strconv.ParseInt("136", 10, 8)
pagesize,_ := strconv.ParseInt("136", 10, 8)
muted := true
deaf := true

var result string
result,_ = conference.CreateListParticipantJSON(conferenceSid, page, pagesize, muted, deaf)

```





### <a name="create_list_participant_xml"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListParticipantXML") CreateListParticipantXML

> List Participant XML

```go
func (me *CONFERENCE_IMPL) CreateListParticipantXML(
            conferenceSid string,
            page *int64,
            pagesize *int64,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| page |  ``` Optional ```  | TODO: Add a parameter description |
| pagesize |  ``` Optional ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
conferenceSid := "ConferenceSid"
page,_ := strconv.ParseInt("136", 10, 8)
pagesize,_ := strconv.ParseInt("136", 10, 8)
muted := true
deaf := true

var result string
result,_ = conference.CreateListParticipantXML(conferenceSid, page, pagesize, muted, deaf)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="phonenumber_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg


#### Get instance
Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.
```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

### <a name="create_view_number_details_json"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateViewNumberDetailsJSON") CreateViewNumberDetailsJSON

> Get Phone Number Details JSON

```go
func (me *PHONENUMBER_IMPL) CreateViewNumberDetailsJSON(phoneNumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Get Phone number Detail |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateViewNumberDetailsJSON(phoneNumber)

```





### <a name="create_view_phone_number_xml"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateViewPhoneNumberXML") CreateViewPhoneNumberXML

> Get Phone Number Details XML

```go
func (me *PHONENUMBER_IMPL) CreateViewPhoneNumberXML(phoneNumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Get Phone number Detail |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateViewPhoneNumberXML(phoneNumber)

```





### <a name="create_buy_number_json"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateBuyNumberJSON") CreateBuyNumberJSON

> Buy Phone Number 

```go
func (me *PHONENUMBER_IMPL) CreateBuyNumberJSON(phoneNumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be purchase |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateBuyNumberJSON(phoneNumber)

```





### <a name="create_buy_number_xml"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateBuyNumberXML") CreateBuyNumberXML

> Buy Phone Number 

```go
func (me *PHONENUMBER_IMPL) CreateBuyNumberXML(phoneNumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be purchase |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateBuyNumberXML(phoneNumber)

```





### <a name="create_release_number_json"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateReleaseNumberJson") CreateReleaseNumberJson

> Release number from account JSON

```go
func (me *PHONENUMBER_IMPL) CreateReleaseNumberJson(phoneNumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be relase |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateReleaseNumberJson(phoneNumber)

```





### <a name="create_release_number_xml"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateReleaseNumberXML") CreateReleaseNumberXML

> Release number from account JSON

```go
func (me *PHONENUMBER_IMPL) CreateReleaseNumberXML(phoneNumber string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be release |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateReleaseNumberXML(phoneNumber)

```





### <a name="create_available_phone_number_xml"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateAvailablePhoneNumberXML") CreateAvailablePhoneNumberXML

> Available Phone Number XML

```go
func (me *PHONENUMBER_IMPL) CreateAvailablePhoneNumberXML(
            numberType string,
            areaCode string,
            pageSize *int64)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | Phone Number type either SMS,Voice or ALL |
| areaCode |  ``` Required ```  | Phone Number Areacode |
| pageSize |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
numberType := "NumberType"
areaCode := "AreaCode"
pageSize,_ := strconv.ParseInt("136", 10, 8)

var result string
result,_ = phoneNumber.CreateAvailablePhoneNumberXML(numberType, areaCode, pageSize)

```





### <a name="create_list_number_json"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateListNumberJSON") CreateListNumberJSON

> List Account's Phone number details json

```go
func (me *PHONENUMBER_IMPL) CreateListNumberJSON(
            page *int64,
            pageSize *int64,
            numberType *string,
            friendlyName *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| numberType |  ``` Optional ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("227", 10, 8)
pageSize,_ := strconv.ParseInt("227", 10, 8)
numberType := "NumberType"
friendlyName := "FriendlyName"

var result string
result,_ = phoneNumber.CreateListNumberJSON(page, pageSize, numberType, friendlyName)

```





### <a name="create_list_number_xml"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateListNumberXML") CreateListNumberXML

> List Account's Phone number details XML

```go
func (me *PHONENUMBER_IMPL) CreateListNumberXML(
            page *int64,
            pageSize *int64,
            numberType *string,
            friendlyName *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| numberType |  ``` Optional ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("227", 10, 8)
pageSize,_ := strconv.ParseInt("227", 10, 8)
numberType := "NumberType"
friendlyName := "FriendlyName"

var result string
result,_ = phoneNumber.CreateListNumberXML(page, pageSize, numberType, friendlyName)

```





### <a name="create_available_phone_number_json"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateAvailablePhoneNumberJson") CreateAvailablePhoneNumberJson

> Available Phone Number Json

```go
func (me *PHONENUMBER_IMPL) CreateAvailablePhoneNumberJson(
            numberType string,
            areaCode string,
            pageSize *int64)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | Number type either SMS,Voice or all |
| areaCode |  ``` Required ```  | Phone Number Area Code |
| pageSize |  ``` Optional ```  | Page Size |



#### Example Usage:
```go
numberType := "NumberType"
areaCode := "AreaCode"
pageSize,_ := strconv.ParseInt("227", 10, 8)

var result string
result,_ = phoneNumber.CreateAvailablePhoneNumberJson(numberType, areaCode, pageSize)

```





### <a name="update_phone_number_xml"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.UpdatePhoneNumberXML") UpdatePhoneNumberXML

> Update Phone Number Details XML

```go
func (me *PHONENUMBER_IMPL) UpdatePhoneNumberXML(
            phoneNumber string,
            friendlyName *string,
            voiceUrl *string,
            voiceMethod models_pkg.HttpMethodEnum,
            voiceFallbackUrl *string,
            voiceFallbackMethod models_pkg.HttpMethodEnum,
            hangupCallback *string,
            hangupCallbackMethod models_pkg.HttpMethodEnum,
            heartbeatUrl *string,
            heartbeatMethod models_pkg.HttpMethodEnum,
            smsUrl *string,
            smsMethod models_pkg.HttpMethodEnum,
            smsFallbackUrl *string,
            smsFallbackMethod models_pkg.HttpMethodEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |
| voiceUrl |  ``` Optional ```  | URL requested once the call connects |
| voiceMethod |  ``` Optional ```  | TODO: Add a parameter description |
| voiceFallbackUrl |  ``` Optional ```  | URL requested if the voice URL is not available |
| voiceFallbackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| hangupCallback |  ``` Optional ```  | TODO: Add a parameter description |
| hangupCallbackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| heartbeatUrl |  ``` Optional ```  | URL requested once the call connects |
| heartbeatMethod |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed time |
| smsUrl |  ``` Optional ```  | URL requested when an SMS is received |
| smsMethod |  ``` Optional ```  | TODO: Add a parameter description |
| smsFallbackUrl |  ``` Optional ```  | URL requested once the call connects |
| smsFallbackMethod |  ``` Optional ```  | URL requested if the sms URL is not available |



#### Example Usage:
```go
phoneNumber := "PhoneNumber"
friendlyName := "FriendlyName"
voiceUrl := "VoiceUrl"
voiceMethod := models_pkg.HttpMethod_GET
voiceFallbackUrl := "VoiceFallbackUrl"
voiceFallbackMethod := models_pkg.HttpMethod_GET
hangupCallback := "HangupCallback"
hangupCallbackMethod := models_pkg.HttpMethod_GET
heartbeatUrl := "HeartbeatUrl"
heartbeatMethod := models_pkg.HttpMethod_GET
smsUrl := "SmsUrl"
smsMethod := models_pkg.HttpMethod_GET
smsFallbackUrl := "SmsFallbackUrl"
smsFallbackMethod := models_pkg.HttpMethod_GET

var result string
result,_ = phoneNumber.UpdatePhoneNumberXML(phoneNumber, friendlyName, voiceUrl, voiceMethod, voiceFallbackUrl, voiceFallbackMethod, hangupCallback, hangupCallbackMethod, heartbeatUrl, heartbeatMethod, smsUrl, smsMethod, smsFallbackUrl, smsFallbackMethod)

```





### <a name="update_phone_number_json"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.UpdatePhoneNumberJSON") UpdatePhoneNumberJSON

> Update Phone Number Details JSON

```go
func (me *PHONENUMBER_IMPL) UpdatePhoneNumberJSON(
            phoneNumberr string,
            friendlyName *string,
            voiceUrl *string,
            voiceMethod models_pkg.HttpMethodEnum,
            voiceFallbackUrl *string,
            voiceFallbackMethod models_pkg.HttpMethodEnum,
            hangupCallback *string,
            hangupCallbackMethod models_pkg.HttpMethodEnum,
            heartbeatUrl *string,
            heartbeatMethod models_pkg.HttpMethodEnum,
            smsUrl *string,
            smsMethod models_pkg.HttpMethodEnum,
            smsFallbackUrl *string,
            smsFallbackMethod models_pkg.HttpMethodEnum)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumberr |  ``` Required ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |
| voiceUrl |  ``` Optional ```  | URL requested once the call connects |
| voiceMethod |  ``` Optional ```  | TODO: Add a parameter description |
| voiceFallbackUrl |  ``` Optional ```  | URL requested if the voice URL is not available |
| voiceFallbackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| hangupCallback |  ``` Optional ```  | TODO: Add a parameter description |
| hangupCallbackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| heartbeatUrl |  ``` Optional ```  | URL requested once the call connects |
| heartbeatMethod |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed time |
| smsUrl |  ``` Optional ```  | URL requested when an SMS is received |
| smsMethod |  ``` Optional ```  | TODO: Add a parameter description |
| smsFallbackUrl |  ``` Optional ```  | URL requested once the call connects |
| smsFallbackMethod |  ``` Optional ```  | URL requested if the sms URL is not available |



#### Example Usage:
```go
phoneNumberr := "PhoneNumberr"
friendlyName := "FriendlyName"
voiceUrl := "VoiceUrl"
voiceMethod := models_pkg.HttpMethod_GET
voiceFallbackUrl := "VoiceFallbackUrl"
voiceFallbackMethod := models_pkg.HttpMethod_GET
hangupCallback := "HangupCallback"
hangupCallbackMethod := models_pkg.HttpMethod_GET
heartbeatUrl := "HeartbeatUrl"
heartbeatMethod := models_pkg.HttpMethod_GET
smsUrl := "SmsUrl"
smsMethod := models_pkg.HttpMethod_GET
smsFallbackUrl := "SmsFallbackUrl"
smsFallbackMethod := models_pkg.HttpMethod_GET

var result string
result,_ = phoneNumber.UpdatePhoneNumberJSON(phoneNumberr, friendlyName, voiceUrl, voiceMethod, voiceFallbackUrl, voiceFallbackMethod, hangupCallback, hangupCallbackMethod, heartbeatUrl, heartbeatMethod, smsUrl, smsMethod, smsFallbackUrl, smsFallbackMethod)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="recording_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".recording_pkg") recording_pkg


#### Get instance
Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.
```go
recording := recording_pkg.NewRECORDING()
```

### <a name="create_view_recording_json"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateViewRecordingJSON") CreateViewRecordingJSON

> View Recording JSON

```go
func (me *RECORDING_IMPL) CreateViewRecordingJSON(recordingSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Search Recording sid |



#### Example Usage:
```go
recordingSid := "RecordingSid"

var result string
result,_ = recording.CreateViewRecordingJSON(recordingSid)

```





### <a name="create_view_recording_xml"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateViewRecordingXML") CreateViewRecordingXML

> View Recording XML

```go
func (me *RECORDING_IMPL) CreateViewRecordingXML(recordingSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Search Recording sid |



#### Example Usage:
```go
recordingSid := "RecordingSid"

var result string
result,_ = recording.CreateViewRecordingXML(recordingSid)

```





### <a name="create_delete_recording_json"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateDeleteRecordingJSON") CreateDeleteRecordingJSON

> Delete Recording Record JSON

```go
func (me *RECORDING_IMPL) CreateDeleteRecordingJSON(recordingSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording Sid to be delete |



#### Example Usage:
```go
recordingSid := "RecordingSid"

var result string
result,_ = recording.CreateDeleteRecordingJSON(recordingSid)

```





### <a name="create_delete_recording_xml"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateDeleteRecordingXML") CreateDeleteRecordingXML

> Delete Recording Record XML

```go
func (me *RECORDING_IMPL) CreateDeleteRecordingXML(recordingSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording Sid to be delete |



#### Example Usage:
```go
recordingSid := "RecordingSid"

var result string
result,_ = recording.CreateDeleteRecordingXML(recordingSid)

```





### <a name="create_list_recording_xml"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateListRecordingXML") CreateListRecordingXML

> List Recording XML

```go
func (me *RECORDING_IMPL) CreateListRecordingXML(
            page *int64,
            pageSize *int64,
            dateCreated *string,
            callSid *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | TODO: Add a parameter description |
| pageSize |  ``` Optional ```  | TODO: Add a parameter description |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| callSid |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("227", 10, 8)
pageSize,_ := strconv.ParseInt("227", 10, 8)
dateCreated := "DateCreated"
callSid := "CallSid"

var result string
result,_ = recording.CreateListRecordingXML(page, pageSize, dateCreated, callSid)

```





### <a name="create_list_recording_json"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateListRecordingJSON") CreateListRecordingJSON

> List Recording  JSON

```go
func (me *RECORDING_IMPL) CreateListRecordingJSON(
            page *int64,
            pageSize *int64,
            dateCreated *string,
            callSid *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| callSid |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("227", 10, 8)
pageSize,_ := strconv.ParseInt("227", 10, 8)
dateCreated := "DateCreated"
callSid := "CallSid"

var result string
result,_ = recording.CreateListRecordingJSON(page, pageSize, dateCreated, callSid)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="transcription_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg


#### Get instance
Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.
```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

### <a name="create_view_transcription_json"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateViewTranscriptionJSON") CreateViewTranscriptionJSON

> View Particular Transcription JSON Response

```go
func (me *TRANSCRIPTION_IMPL) CreateViewTranscriptionJSON(transcriptionSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionSid |  ``` Required ```  | Unique Transcription ID |



#### Example Usage:
```go
transcriptionSid := "TranscriptionSid"

var result string
result,_ = transcription.CreateViewTranscriptionJSON(transcriptionSid)

```





### <a name="create_view_transcription_xml"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateViewTranscriptionXML") CreateViewTranscriptionXML

> View Particular Transcription XML Response

```go
func (me *TRANSCRIPTION_IMPL) CreateViewTranscriptionXML(transcriptionSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionSid |  ``` Required ```  | Default Value  Default Value |



#### Example Usage:
```go
transcriptionSid := "TranscriptionSid"

var result string
result,_ = transcription.CreateViewTranscriptionXML(transcriptionSid)

```





### <a name="create_audio_url_transcription_json"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateAudioURLTranscriptionJSON") CreateAudioURLTranscriptionJSON

> Audio URL Transcription JSON

```go
func (me *TRANSCRIPTION_IMPL) CreateAudioURLTranscriptionJSON(audioUrl string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| audioUrl |  ``` Required ```  | Audio url |



#### Example Usage:
```go
audioUrl := "AudioUrl"

var result string
result,_ = transcription.CreateAudioURLTranscriptionJSON(audioUrl)

```





### <a name="create_audio_url_transcription_xml"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateAudioURLTranscriptionXML") CreateAudioURLTranscriptionXML

> Audio URL Transcription XML

```go
func (me *TRANSCRIPTION_IMPL) CreateAudioURLTranscriptionXML(audioUrl string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| audioUrl |  ``` Required ```  | Audio url |



#### Example Usage:
```go
audioUrl := "AudioUrl"

var result string
result,_ = transcription.CreateAudioURLTranscriptionXML(audioUrl)

```





### <a name="create_recording_transcription_json"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateRecordingTranscriptionJSON") CreateRecordingTranscriptionJSON

> Recording Transcription JSON

```go
func (me *TRANSCRIPTION_IMPL) CreateRecordingTranscriptionJSON(recordingSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording sid |



#### Example Usage:
```go
recordingSid := "RecordingSid"

var result string
result,_ = transcription.CreateRecordingTranscriptionJSON(recordingSid)

```





### <a name="create_recording_transcription_xml"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateRecordingTranscriptionXML") CreateRecordingTranscriptionXML

> Recording Transcription XML

```go
func (me *TRANSCRIPTION_IMPL) CreateRecordingTranscriptionXML(recordingSid string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording sid |



#### Example Usage:
```go
recordingSid := "RecordingSid"

var result string
result,_ = transcription.CreateRecordingTranscriptionXML(recordingSid)

```





### <a name="create_list_transcription_json"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateListTranscriptionJSON") CreateListTranscriptionJSON

> Get All transcription in JSON response

```go
func (me *TRANSCRIPTION_IMPL) CreateListTranscriptionJSON(
            page *int64,
            pageSize *int64,
            status models_pkg.StatusEnum,
            dateTranscribed *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | TODO: Add a parameter description |
| pageSize |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateTranscribed |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("227", 10, 8)
pageSize,_ := strconv.ParseInt("227", 10, 8)
status := models_pkg.Status_INPROGRESS
dateTranscribed := "DateTranscribed"

var result string
result,_ = transcription.CreateListTranscriptionJSON(page, pageSize, status, dateTranscribed)

```





### <a name="create_list_transcription_xml"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateListTranscriptionXML") CreateListTranscriptionXML

> Get All transcription in XML response

```go
func (me *TRANSCRIPTION_IMPL) CreateListTranscriptionXML(
            page *int64,
            pageSize *int64,
            status models_pkg.StatusEnum,
            dateTranscribed *string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | TODO: Add a parameter description |
| pageSize |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateTranscribed |  ``` Optional ```  | TODO: Add a parameter description |



#### Example Usage:
```go
page,_ := strconv.ParseInt("227", 10, 8)
pageSize,_ := strconv.ParseInt("227", 10, 8)
status := models_pkg.Status_INPROGRESS
dateTranscribed := "DateTranscribed"

var result string
result,_ = transcription.CreateListTranscriptionXML(page, pageSize, status, dateTranscribed)

```





[Back to List of Controllers](#list_of_controllers)
## <a name="usage_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".usage_pkg") usage_pkg


#### Get instance
Factory for the ``` USAGE ``` interface can be accessed from the package usage_pkg.
```go
usage := usage_pkg.NewUSAGE()
```

### <a name="create_list_usage"></a>![Method: ](http://apidocs.io/img/method.png ".usage_pkg.CreateListUsage") CreateListUsage

> Get all usage 

```go
func (me *USAGE_IMPL) CreateListUsage(
            productCode string,
            startDate string,
            endDate string)(string,error)
```

#### Parameters: 

| Parameter | Tags | Description |
|-----------|------|-------------|
| productCode |  ``` Required ```  ``` DefaultValue ```  | Product Code |
| startDate |  ``` Required ```  ``` DefaultValue ```  | Start Usage Date |
| endDate |  ``` Required ```  ``` DefaultValue ```  | End Usage Date |



#### Example Usage:
```go
productCode := "0"
startDate := "2016-09-06"
endDate := "2016-09-06"

var result string
result,_ = usage.CreateListUsage(productCode, startDate, endDate)

```





[Back to List of Controllers](#list_of_controllers)


