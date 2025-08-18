import { useEffect, useRef, useState } from "react";
import StarRating from "./StarRating";
import Loader from "./Loader";
import type { TMovie } from "../types";
import { key } from "../Hooks/useMovie";

type ExMovie = {
  Title: string;
  Poster: string;
  Runtime: string;
  Year: string;
  imdbRating: string;
  Plot: string;
  Released: string;
  Actors: string;
  Director: string;
  Genre: string;
};

type SelectedMovieProps = {
  onCloseMovie: () => void;
  onAddWatched: (movie: TMovie) => void;
  watchedMovies: TMovie[];
  selectedID: string;
};

const MovieDetails = ({
  selectedID,
  onCloseMovie,
  onAddWatched,
  watchedMovies,
}: SelectedMovieProps) => {
  const [movie, setMovie] = useState<ExMovie | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [userRating, setUserRating] = useState<number>(0);

  const countRef = useRef(0);

  useEffect(() => {
    if (userRating) countRef.current++;
  }, [userRating]);

  const isWatched = watchedMovies
    .map((movie) => movie.imdbID)
    .includes(selectedID);

  const watchedUserRating = watchedMovies.find(
    (movie) => movie.imdbID === selectedID,
  )?.userRating;

  // Get movie details
  useEffect(
    function () {
      const controller = new AbortController();

      const getMovieDetails = async () => {
        setIsLoading(true);
        const queryUrl = `http://www.omdbapi.com/?apikey=${key}&i=${selectedID}`;

        const init: RequestInit = {
          signal: controller.signal,
        };

        const res = await fetch(queryUrl, init);

        const data = await res.json();
        setMovie(data);
        setIsLoading(false);
      };

      getMovieDetails();

      return function () {
        controller.abort();
      };
    },
    [selectedID],
  );

  // Set title from movie details
  useEffect(
    function () {
      if (movie) {
        document.title = `Movie - ${movie.Title}`;

        return function () {
          document.title = "usePopcorn";
        };
      }
    },
    [movie],
  );

  if (isLoading || !movie) return <Loader />;

  const {
    Title: title,
    Poster: poster,
    Year: year,
    Runtime: runtime,
    imdbRating,
    Plot: plot,
    Released: released,
    Actors: actors,
    Director: director,
    Genre: genre,
  } = movie;

  const handleAdd = () => {
    if (!movie) return;

    const watchedMovie: TMovie = {
      imdbID: selectedID,
      title,
      year,
      poster,
      runtime: Number(runtime.split(" ").at(0)),
      userRating,
      imdbRating: Number(imdbRating),
      countRatingDecisions: countRef.current,
    };

    onAddWatched(watchedMovie);
    onCloseMovie();
  };

  return (
    <div className="details">
      {isLoading || !movie ? (
        <Loader />
      ) : (
        <>
          <header>
            <button className="btn-back" onClick={onCloseMovie}>
              &larr;{" "}
            </button>
            <img src={poster} alt={`poster of ${movie}`} />
            <div className="details-overview">
              <h2>{title}</h2>
              <p>
                {released} &bull; {runtime}
              </p>
              <p>{genre}</p>
              <p>
                <span>⭐️</span>
                {imdbRating} IMDB Rating
              </p>
            </div>
          </header>

          <section>
            <div className="rating">
              {!isWatched ? (
                <>
                  <StarRating
                    maxRating={10}
                    onExternalSetRating={setUserRating}
                    size={24}
                  />
                  {userRating > 0 && (
                    <button className="btn-add" onClick={handleAdd}>
                      + Add to list
                    </button>
                  )}
                </>
              ) : (
                <>
                  <p>You rated this movie: {watchedUserRating} ⭐️</p>
                </>
              )}
            </div>
            <p>
              <em>{plot}</em>
            </p>
            <p>Starring {actors}</p>
            <p>Directed by {director}</p>
          </section>
        </>
      )}
    </div>
  );
};

export default MovieDetails;
