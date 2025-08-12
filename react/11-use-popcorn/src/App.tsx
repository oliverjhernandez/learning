import Nav from "./Components/Nav";
import Main from "./Components/Main";
import "./index.css";
import { useEffect, useState, type JSX } from "react";
import type { TMovie } from "./types";
import NumResults from "./Components/NumResults";
import Search from "./Components/Search";
import Logo from "./Components/Logo";
import MovieList from "./Components/MovieList";
import Box from "./Components/Box";
import WatchedSummary from "./Components/WatchedSummary";
import WatchedMovieList from "./Components/WatchedMovieList";
import Loader from "./Components/Loader";
import ErrorMessage from "./Components/ErrorMessage";
import MovieDetails from "./Components/MovieDetails";

export const key = "644cebbb";

const App = (): JSX.Element => {
  const [query, setQuery] = useState<string>("");
  const [movies, setMovies] = useState<TMovie[]>([]);
  const [watched, setWatched] = useState<TMovie[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [err, setError] = useState<string>("");
  const [selectedID, setSelectedID] = useState<string>("");

  const handleSelectMovie = (id: string) => {
    setSelectedID(selectedID === id ? "" : id);
  };

  const handleCloseMovie = () => {
    setSelectedID("");
  };

  const handleAddWatched = (movie: TMovie) => {
    setWatched((watched) => [...watched, movie]);
  };

  const handleDeleteWatched = (id: string) => {
    setWatched((watched) => watched.filter((movie) => movie.imdbID !== id));
  };

  useEffect(
    function () {
      localStorage.setItem("watched", JSON.stringify(watched));
    },
    [watched],
  );

  useEffect(
    function () {
      async function fetchMovies() {
        try {
          const queryUrl = `http://www.omdbapi.com/?apikey=${key}&s=${query}`;
          setIsLoading(true);
          setError("");
          const res = await fetch(queryUrl);

          if (!res.ok)
            throw new Error("Something went wrong while fetching movies");

          const data = await res.json();
          if (data.Response === "False") {
            throw new Error("Movie not found");
          }

          const movies: TMovie[] = data.Search.map((movie: any) => ({
            imdbID: movie.imdbID,
            title: movie.Title,
            year: movie.Year,
            poster: movie.Poster,
          }));

          setMovies(movies);
          setError("");
        } catch (err) {
          if (err instanceof Error) {
            console.log(err.message);
            if (err.name !== "AbortError") {
              setError(err.message);
            }
          }
        } finally {
          setIsLoading(false);
        }
      }

      if (query.length <= 3) {
        setMovies([]);
        setError("");
        return;
      }

      fetchMovies();
    },
    [query],
  );

  return (
    <div>
      <Nav>
        <Logo />
        <Search query={query} setQuery={setQuery} />
        <NumResults results={movies.length} />
      </Nav>
      <Main>
        <Box>
          {isLoading && <Loader />}
          {!isLoading && err && <ErrorMessage message={err} />}
          {!isLoading && !err && (
            <MovieList movies={movies} onSelectMovie={handleSelectMovie} />
          )}
        </Box>
        <Box>
          {selectedID ? (
            <MovieDetails
              selectedID={selectedID}
              onCloseMovie={handleCloseMovie}
              onAddWatched={handleAddWatched}
              watchedMovies={watched}
            />
          ) : (
            <>
              <WatchedSummary watched={watched} />
              <WatchedMovieList
                watched={watched}
                onDelete={handleDeleteWatched}
              />
            </>
          )}
        </Box>
      </Main>
    </div>
  );
};

export default App;
