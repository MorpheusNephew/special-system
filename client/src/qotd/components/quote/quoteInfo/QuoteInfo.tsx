import React from 'react';

type QuoteInfoParams = {
    quote?: string,
    language?: string,
    likes?: number;
    author?: string;
    tags?: string[];
}

const QuoteInfo: React.FC<QuoteInfoParams> = ({quote, language, likes, author, tags}) => (
    <>
      <div>Quote: {quote}</div>
      <div>Language: {language}</div>
      <div>Likes: {likes}</div>
      <div>Author: {author}</div>
      <div>Tags: {tags?.join(", ")}</div>
    </>
);

export { QuoteInfo }