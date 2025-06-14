export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {/* <header></header> */}
        <h1>Welcome to ASU Module</h1>
        {children}
        {/* <footer></footer> */}
      </body>
    </html>
  );
}