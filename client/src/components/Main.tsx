import { Component, createSignal } from 'solid-js';
import { generateImages } from '../controllers/image';
import { GenerateImagesReq } from '../models/models';
import Images from './Images';
import Nav from './Nav';
import Search from './Search';

const Main: Component = () => {
  const [isLoading, setIsLoading] = createSignal(false);
  const [images, setImages] = createSignal([]);

  const onSearch = async (input: string) => {
    setIsLoading(true);
    const res = await generateImages({ theme: input } as GenerateImagesReq);

    if (!res.data) {
      // TODO: add error modal
    }

    setImages(res.data.urls);
    setIsLoading(false);
  };

  const displayImages = () => {
    if (isLoading()) {
      return <p class="text-xl text-slate-500">loading...</p>;
    } else if (images().length > 0) {
      return <Images urls={images()} />;
    } else {
      return (
        <p class="text-xl text-slate-500 px-16">
          press 'search' to generate a moodboard. if a theme is not provided, a
          random moodboard will be generated.
        </p>
      );
    }
  };

  return (
    <div class="bg-white w-screen min-h-screen flex flex-col items-center gap-y-14 md:gap-y-20 2xl:gap-y-24 pb-20">
      <Nav />
      <Search onSearch={onSearch} />
      {displayImages()}
    </div>
  );
};

export default Main;
