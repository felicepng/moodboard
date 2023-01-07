import type { Component } from 'solid-js';
import { images, theme } from '../App';

const Text: Component = () => {
  return (
    <>
      {images.loading && (
        <p class="text-center text-xl md:text-2xl text-slate-500">loading...</p>
      )}
      {!images.loading && images()?.message && (
        <p class="text-center text-xl md:text-2xl text-red-500">
          Status {images().status}: {images().message}
        </p>
      )}
      {!images.loading && !theme()[0] && (
        <p class="text-center text-xl md:text-2xl text-slate-500 px-14 md:px-20">
          press 'search' to generate a moodboard. if a theme is not provided, a
          random moodboard will be generated.
        </p>
      )}
    </>
  );
};

export default Text;
