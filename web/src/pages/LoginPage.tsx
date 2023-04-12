import React, { useState } from "react";
import { Avatar, Box, Typography } from "@mui/material";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import { LoginForm, User } from "../components/LoginForm";

export default function LoginPage() {
  const [user, setUser] = useState<User | null>(null);

  return (
    <Box
      sx={{
        marginTop: 8,
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      {!!user ? (
        <Typography component="h1" variant="h5">
          Hello, {user.username}
        </Typography>
      ) : (
        <>
          <Avatar sx={{ m: 1, bgcolor: "primary.main" }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <LoginForm onSuccess={setUser} />
        </>
      )}
    </Box>
  );
}
