import { screen } from "@testing-library/react";
import { runContract } from "./cdc";

runContract("home/no_movies_found.yaml", (_) => {
  expect(screen.getByText(/Create Movie/i)).toBeInTheDocument()
});

runContract("home/multiple_movies_found.yaml", (app, _) => {
  expect(screen.getByText(/Create Movie/i)).toBeInTheDocument()

  app.textShouldExist.forEach((text, _) => {
    expect(screen.getByText(text)).toBeInTheDocument()
  });
});
