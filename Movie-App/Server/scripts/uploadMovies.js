const fs = require("fs").promises;
const path = require("path");
const { Client } = require("pg");

async function uploadMovies() {
  try {
    const client = new Client({
      user: "postgres",
      database: "movies",
      password: "tanzeem",
    });

    await client.connect();

    const data = await fs.readFile(path.resolve(__dirname, "../data/movies.json"), "utf-8");
    const movies = JSON.parse(data);

    const promises = [];
    for (const movie of movies) {
      promises.push(
        client.query(
          `
          INSERT INTO movies (
            adult, backdrop_path, movie_id, title, original_language, original_title,
            overview, poster_path, media_type, genre_ids, popularity, release_date,
            video, vote_average, vote_count
          )
          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
          `,
          [
            movie.adult,
            movie.backdrop_path,
            movie.id,
            movie.title,
            movie.original_language,
            movie.original_title,
            movie.overview,
            movie.poster_path,
            movie.media_type,
            movie.genre_ids,
            movie.popularity,
            movie.release_date,
            movie.video,
            movie.vote_average,
            movie.vote_count,
          ]
        )
      );
    }

    await Promise.all(promises);
    console.log("Movies uploaded");
    await client.end();
  } catch (err) {
    console.error(err);
  }
}

uploadMovies();
