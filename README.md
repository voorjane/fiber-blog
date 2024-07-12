# Blog page made with Fiber

## About this project

This project is a blog page aimed at learning fiber web-framework

Technologies used:

* [Golang](https://go.dev/)
* [Fiber](https://github.com/gofiber)
* [PostgreSQL](https://www.postgresql.org/)
* [GORM](https://gorm.io)
* [REST](https://ru.wikipedia.org/wiki/REST)
* [Docker](https://www.docker.com/)

## Getting started

This is an example of how you may run this project locally. Follow these steps:

1. Clone the repo
`git clone https://github.com/voorjane/fiber-blog`

2. Make sure [docker](https://www.docker.com/) is installed

3. Run
`make`

4. Go to `localhost:8888` in your browser

## Usage

As you run this project, there are no articles yet. To create one, go to [Admin panel](localhost:8888/sign_in) and log in
- login: `admin`
- password: `admin`

Then, simply type title, announce and article text in matching fields, and click `Create`

Article text supports Markdown, so check it out!

This page has a rate limiter of 100 RPS, you can check it running `rate_test.sh` script

## TODO

- .env instead of exposed postgres credentials
- Add `Remember me` function so you can use admin panel more then one time

## Contact

Alexey Vasilchenko - [@ny\_tbl\_alesha](t.me/ny_tbl_alesha) - [vasilchenkoaleksey@gmail.com](mailto:vasilchenkoaleksey@gmail.com)

