import "./App.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ConfigProvider, ThemeConfig } from "antd";
import Home from "./pages/home";
import { GoogleOAuthProvider } from "@react-oauth/google";
import React from "react";
// import { theme as antdTheme } from "antd";

const queryClient = new QueryClient();
const theme: ThemeConfig = {
  components: {
    Layout: {
      siderBg: "white",
    },
  },
};

function App() {
  return (
    <GoogleOAuthProvider clientId={import.meta.env.VITE_GOOGLE_CLIENT_ID}>
      <QueryClientProvider client={queryClient}>
        <ConfigProvider theme={theme}>
          <React.StrictMode>
            <Home />
          </React.StrictMode>
        </ConfigProvider>
      </QueryClientProvider>
    </GoogleOAuthProvider>
  );
}

export default App;
