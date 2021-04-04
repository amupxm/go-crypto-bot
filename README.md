# Golang cryptocurrency telegram bot

simple solution to get latest crypto price in any group/channel.

## Usage

use sample.env file to create an . env file which will contains your telegram bot token. then build the app and run as binary.

```bash
sed -i 's/telegram bot token/TOKEN/g' sample.env 
mv sample.env .env
go install . 
go build . 
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)