React Native

npx create-expo-app --template blank <project-name>

Download expo go on your smartphone

npm start

Scan QR code on your smartphone

Enable fast refresh by shaking your smartphone if disabled


Publishing React Native Apps 
Configure the project -> Build App binaries with expo cloud services using EAS (Expo Application Services)-> Submit to App Store

Install the EAS CLI -> npm install -g eas-cli -> eas login
Configure the project -> eas build:configure
It will create a eas.json file. 

Follow the documentation for building the android APK (https://docs.expo.dev/build-reference/apk/)

Change buildType to ask in eas.json file
Run this command: eas build -p android --profile preview

Once the build is done, we will get a link to the APK file, which we can then download.
