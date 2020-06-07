import React, { useEffect, useState } from "react";
import Paper from "@material-ui/core/Paper";
import { getQuote } from "../../resources";
import { QuoteResponse } from "../../models/quote";
import { QuoteInfo } from "./quoteInfo/QuoteInfo";
import { makeStyles } from "@material-ui/core";

const useStyles = makeStyles({
  root: {
    marginTop: "5%",
  },
});

const Quote = () => {
  const [quoteResponse, setQuoteResponse] = useState<QuoteResponse | undefined>(
    undefined
  );

  useEffect(() => {
    getQuote().then((response) => {
      setQuoteResponse(response.data);
    });
  }, []);

  const styles = useStyles();

  return (
    <Paper className={styles.root}>
      {quoteResponse && (
        <QuoteInfo
          author={quoteResponse.author}
          language={quoteResponse.language}
          likes={quoteResponse.likes}
          quote={quoteResponse.quote}
          tags={quoteResponse.tags}
        />
      )}
    </Paper>
  );
};

export { Quote };
