module.exports = (api) => {
  // cache babel configurations via NODE_ENV environments
  // development, production, testing, staging...etc
  api.cache(() => process.env.NODE_ENV);

  return {
    presets: ["next/babel"],
    plugins: [
      // leverage the 'babel-plugin-styled-component' plugin for...
      // 1.) consistently hashed component classNames between environments (a must for server-side rendering)
      // 2.) better debugging through automatic annotation of your styled components based on their context in the file system, etc.
      // 3.) minification for styles and tagged template literals styled-components usages
      [
        "styled-components",
        {
          ssr: true,
          displayName: true,
          preprocess: false,
        },
      ],
    ],
  };
};
