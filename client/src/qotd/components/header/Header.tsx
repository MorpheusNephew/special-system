import * as React from "react";
import AppBar from "@material-ui/core/AppBar";
import Typography from "@material-ui/core/Typography";

const Header = () => (
  <AppBar>
    <Typography variant="h6" align="center">
      Welcome to Quote of the Day app
    </Typography>
  </AppBar>
);

export { Header };
