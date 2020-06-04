import React, { useEffect, useState } from "react";
import { getQuote } from "../../resources";
import { QuoteResponse } from "../../models/quote";
import { QuoteInfo } from "./quoteInfo/QuoteInfo";

const Quote = () => {
  const [quoteResponse, setQuoteResponse] = useState<QuoteResponse | undefined>(
    undefined
  );

  useEffect(() => {
    getQuote().then((response) => {
      setQuoteResponse(response.data);
    });
  }, []);

  return (
    <>
      {quoteResponse && (
        <QuoteInfo
          author={quoteResponse.author}
          language={quoteResponse.language}
          likes={quoteResponse.likes}
          quote={quoteResponse.quote}
          tags={quoteResponse.tags}
        />
      )}
    </>
  );
};

export { Quote };
