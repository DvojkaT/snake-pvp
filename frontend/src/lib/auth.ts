export function getAuth() {
  let uuid = localStorage.getItem("uuid");
  if (!uuid) {
    uuid = crypto.randomUUID()
    localStorage.setItem("uuid", uuid);
  }
  return uuid;
}

export function getName() {
  let name = localStorage.getItem("name");
  if (!name) {
    name = "player-"+Math.floor(Math.random() * 1000) // до 1000
    localStorage.setItem("name", name);
  }

  return name;
}

export function setName(name: string) {
  localStorage.setItem("name", name);
}
