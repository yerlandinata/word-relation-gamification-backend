# Building Indonesian Hyponym-Hypernym Semantic Relations Corpus and Model Using Pattern-Based, Crowdsourcing, and Machine Learning Approach
This repository stores the *backend web service* code for the crowdsourcing application used in the research

## How to run
1. Clone
2. Create config files, `.env`
```
DB_HOST=
DB_NAME=
DB_USER=
DB_PASS=
TARGET_ANNOTATION_COUNT_PER_WORD_PAIR=
GAME_TIME_LIMIT_MS=time limit for a game level
SECRET=jwt secret
NOTSURE_WRT_ID=4
PORT=
```
3. Set up the PostgreSQL DB
```
execute the SQL scripts in db/schema folder in correct order.
Do not use the seed in the db/seed folder, just take example. Note that the seed in this repo not necessarily use the latest DB schema.
```
4. Use the build script `build.sh` to build the application
5. Start the web service with `./gamification` or use *systemd*
6. Make sure the frontend application is running (repo: https://github.com/yerlandinata/word-relation-gamification-frontend)
7. Eliminate annotators with `./gamification filter-annotators`

## Results
The crowdsourced corpus can be viewed here: https://github.com/yerlandinata/korpus-hiponim-hipernim

## App Screenshots
![screenshots](https://github.com/yerlandinata/word-relation-gamification-frontend/blob/master/app_chosen_ss.png)

## Citation
*short paper coming soon*

Erlandinata, Y. (2020). Pembangunan Korpus dan Model Relasi Semantik Hiponim-Hipernim Bahasa Indonesia dengan Pendekatan Pattern-Based, Crowdsourcing, dan Machine Learning (Skripsi). Universitas Indonesia, Depok, Indonesia.
