export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <title>asu</title>
      </head>
      <body>
        {/* <header></header> */}
        <h1>Welcome to ASU</h1>
        {children}
        {/* <footer></footer> */}
      </body>
    </html>
  );
}