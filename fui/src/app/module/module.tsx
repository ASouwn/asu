export function Module({ username, content }: {
  username: string;
  content: string;
}) {
  return (
    <div>
        <h1>{username}</h1>
        <p>{content}</p>
    </div>
  );
}