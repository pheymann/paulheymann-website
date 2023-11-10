import { useState } from "react";
import { useNavigate } from 'react-router-dom';

export default function CreateMovie({ fetch, token }) {
  const [title, setTitle]             = useState("");
  const [releaseYear, setReleaseYear] = useState(0);
  const navigate = useNavigate();

  const handleClick = async () => {
    const id = `${title}::${Date.now()}`;

    try {
      await fetch.apply('/api/movie', {
        method: 'PUT',
        headers: {
          'Authorization': token,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          id: id,
          title: title,
          releaseYear: parseFloat(releaseYear)
        })
      })

      navigate("/");
    } catch (e) {
      console.error(e);
    }
  }

  return(
    <div>
      <input  value={ title }
              placeholder='Title'
              autoFocus={ true }
              onChange={ e => setTitle(e.target.value) }
      />
      <br/>
      <input  value={ releaseYear }
              placeholder='Release Year'
              onChange={ e => setReleaseYear(e.target.value) }
      />
      <br/>
      <button onClick={ _ => handleClick() }>
        Create
      </button>
    </div>
  );
}
