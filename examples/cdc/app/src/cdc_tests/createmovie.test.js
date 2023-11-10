import { fireEvent, screen, waitFor } from "@testing-library/react";
import userEvent from '@testing-library/user-event';
import { runContract } from "./cdc";

runContract("create_movie/unauthorized.yaml",
  (_) => {
    expect(screen.queryByRole('link', { name : 'Create Movie' })).not.toBeInTheDocument()
  },
  async () => {
    // path: /create
    const createButton = await waitFor(() => screen.findByRole('button', { name: 'Create' }));
    await fireEvent.change(screen.getByPlaceholderText(/Title/i), { target: { value: 'test' } });
    await fireEvent.change(screen.getByPlaceholderText(/Release Year/i), { target: { value: 2023 } });
    await userEvent.click(createButton);
  }
);

runContract("create_movie/created.yaml",
  (app, _) => {
    expect(screen.getByText(/Create Movie/i)).toBeInTheDocument()

    app.textShouldExist.forEach(text => {
      expect(screen.getByText(text)).toBeInTheDocument()
    });
  },
  async () => {
    // path: /
    const navigateButton = await waitFor(() => screen.findByRole('link', { name: 'Create Movie' }));
    await userEvent.click(navigateButton);

    // path: /create
    const createButton = await waitFor(() => screen.findByRole('button', { name: 'Create' }));
    await fireEvent.change(screen.getByPlaceholderText(/Title/i), { target: { value: 'test' } });
    await fireEvent.change(screen.getByPlaceholderText(/Release Year/i), { target: { value: 2023 } });
    await userEvent.click(createButton);
  }
);
