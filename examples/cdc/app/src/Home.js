import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

export default function Home({ fetch }) {
  const [movies, setMovies] = useState([]);

  useEffect(() => {
    const fetchMovies = async () => {
      try {
        await fetch.apply('/api/movie/all', {
          method: 'GET'
        })
          .then(res => res.json())
          .then(data => setMovies(data));
      } catch (e) {
        console.error(e);
      }
    };

    fetchMovies();
  }, [fetch]);

  return(
    <div>
      <ul>
        { movies.map(movie => <li key={movie.id}>{movie.title}</li>) }
      </ul>

      <Link to='/create'>
        Create Movie
      </Link>
    </div>
  );
}
