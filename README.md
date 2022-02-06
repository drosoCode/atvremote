# Android TV remote

## About this project
This tool implements the remote control protocol for Android TV and Google TV.

The v1 is for the old [Android TV Remote Control](https://www.apkmirror.com/apk/google-inc/remote-control/) app <br>
The v2 is for the newer [Google TV](https://play.google.com/store/apps/details?id=com.google.android.videos) app used when your android tv's [remote services](https://www.apkmirror.com/apk/google-inc/android-tv-remote-service-android-tv/) version is >= 5 <br>
The [anymote protocol](https://code.google.com/archive/p/anymote-protocol/) (old protocol for google tv < 2014) is not supported.

## Quickstart
Check on your android tv device the `Remote Services` version (in Settings > Apps > See all apps > Android TV Remote Service).
If the version is >= 5, you sould use the v2 protocol.

First, start the pairing process `./atvremote -ip="192.168.1.20" -version=2 -pair`<br>
Then, you can send a list of keypress like this (this sould open the settings) `./atvremote -ip="192.168.1.20" -version=2 command="HOME;HOME;UP;RIGHT;RIGHT;ENTER`


## Usage

You can use this project as a library (see the godoc page) or directly in cli (see the releases page).
In CLI mode, you must at least specify the IP to your android tv box.
Here is the list of the supported flags:

|Name|Required|Default|Description|Example|
|----|--------|-------|-----------|-------|
|pair|No|False|Pairing mode (if not specified the tool will start in command mode and assume that the pairing have already been done)|`-pair`
|ip|Yes|None|IP or hostname of your android tv device| `-ip="192.168.1.20"`|
|cert|No|`./cert.pem`|Path to your certificate file (will be created in pairing mode)| `-cert="./cert.pem"`|
|key|No|`./key.pem`|Path to your private key file (will be created in pairing mode)| `-key="./key.pem"`|
|version|No|`1`|Protocol version| `-version=1`|
|command|No|None|List of keypress to emulate (see below)| `-command="HOME;HOME;UP;RIGHT;RIGHT;ENTER"`|
|open|No|None|Open an app using an intent (V1 Only)|`-open="com.netflix.ninja/.MainActivity"`|
|link|No|None|Open a link (V2 Only)|`-open="https://netflix.com/title/"`|
|info|No|None|Print a json object of the device's infos (V2 Only)|`-info`|

### Additional information
 - Command flag: a string with alphanumeric characters will automaticaly be converted to keypresses. You can also send some specific keypresses defined in `pkg/common/types.go` prefixed with `KEYCODE_` and separated with semicolons. There are some shortcuts for the most common keys (UP, DOWN, LEFT, RIGHT, HOME, BACK, VOLP, VOLM, MUTE, PWR, ENTER).
 For examle this string `a;UP;bc;KEYCODE_TV_INPUT_HDMI_1` will be parsed as `[RemoteKeyCode_KEYCODE_A, RemoteKeyCode_KEYCODE_DPAD_UP, RemoteKeyCode_KEYCODE_B, RemoteKeyCode_KEYCODE_C, KEYCODE_TV_INPUT_HDMI_1]` (see `pkg/common/parsekey.go` for more info).
 - Open flag: an activity in in the format `yourpackagename/.activityname`.
 - Link flag: You can start an app using a deeplink defined in its appmanifest (like described [here](https://developer.android.com/training/app-links/deep-linking)). By default, if no deeplink is found, the link will open in a browser.


## Acknowledgments
This project wouldn't have been possible without these awesome projects which reverse-engineered these protocols.
 - [Aymkdn](https://github.com/Aymkdn)'s wiki on the protocols [V1](https://github.com/Aymkdn/assistant-freebox-cloud/wiki/Google-TV-(aka-Android-TV)-Remote-Control) and [V2](https://github.com/Aymkdn/assistant-freebox-cloud/wiki/Google-TV-(aka-Android-TV)-Remote-Control-(v2))
 - [louis49](https://github.com/louis49/androidtv-remote)'s [androidtv-remote](https://github.com/louis49/androidtv-remote) js implementation (especially for the v2 proto files)