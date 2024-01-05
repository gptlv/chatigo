"use client";
import "./globals.css";
import AuthContextProvider from "@/modules/AuthProvider";
import WebSocketProvider from "@/modules/WebSocketProvider";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <AuthContextProvider>
        <WebSocketProvider>
          <html lang="en">
            <body>{children}</body>
          </html>
        </WebSocketProvider>
      </AuthContextProvider>
    </>
  );
}
