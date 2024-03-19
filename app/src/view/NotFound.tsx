import notFoundImage from '../assets/images/notFound.png';

const NotFound = () => {
  return (
    <div className='not-found-container'>
      <img
        src={notFoundImage}
        alt='404 Not Found'
        className='not-found-image'
      />
      <h1>404 - Page Not Found</h1>
      <p>The page you are looking for does not exist.</p>
    </div>
  );
};

export default NotFound;
