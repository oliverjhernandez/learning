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
import { useMovies } from "./Hooks/useMovie";

const App = (): JSX.Element => {
  const [query, setQuery] = useState<string>("");
  const [watched, setWatched] = useState<TMovie[]>([]);
  const [selectedID, setSelectedID] = useState<string>("");

  const { movies, isLoading, err } = useMovies(query);

  function handleSelectMovie(id: string) {
    setSelectedID(selectedID === id ? "" : id);
  }

  function handleCloseMovie() {
    setSelectedID("");
  }

  function handleAddWatched(movie: TMovie) {
    setWatched((watched) => [...watched, movie]);
  }

  function handleDeleteWatched(id: string) {
    setWatched((watched) => watched.filter((movie) => movie.imdbID !== id));
  }

  useEffect(
    function () {
      localStorage.setItem("watched", JSON.stringify(watched));
    },
    [watched],
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
