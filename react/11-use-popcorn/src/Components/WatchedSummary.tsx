import type { TMovie } from "../types";

type SummaryProps = {
  watched: TMovie[];
};

const average = (arr: number[]) =>
  arr.reduce((acc, cur, _, arr) => acc + cur / arr.length, 0);

const WatchedSummary = (props: SummaryProps) => {
  const validImdbRatings = props.watched.filter(
    (movie): movie is TMovie & { imdbRating: number } =>
      typeof movie.imdbRating === "number" && !isNaN(movie.imdbRating),
  );
  const avgImdbRating = average(
    validImdbRatings.map((movie) => movie.imdbRating),
  );

  const validUserRatings = props.watched.filter(
    (movie): movie is TMovie & { userRating: number } =>
      typeof movie.userRating === "number" && !isNaN(movie.userRating),
  );
  const avgUserRating = average(
    validUserRatings.map((movie) => movie.userRating),
  );

  const validRuntime = props.watched.filter(
    (movie): movie is TMovie & { runtime: number } =>
      typeof movie.runtime === "number" && !isNaN(movie.runtime),
  );
  const avgRuntime = average(validRuntime.map((movie) => movie.runtime));

  return (
    <div className="summary">
      <h2>Movies you watched</h2>
      <div>
        <p>
          <span>#️⃣</span>
          <span>{props.watched.length} movies</span>
        </p>
        <p>
          <span>⭐️</span>
          <span>{avgImdbRating}</span>
        </p>
        <p>
          <span>🌟</span>
          <span>{avgUserRating}</span>
        </p>
        <p>
          <span>⏳</span>
          <span>{avgRuntime} min</span>
        </p>
      </div>
    </div>
  );
};

export default WatchedSummary;
