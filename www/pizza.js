import {
  html,
  Component,
  render
} from 'https://unpkg.com/htm/preact/standalone.mjs';
const toRad = num => {
  return (num * Math.PI) / 180;
};
const toDeg = num => {
  return (num * 180) / Math.PI;
};

const distanceFrom = (from, to) => {
  const R = 6371; // km
  const dLat = toRad(from.latitude - to.latitude);
  const dLon = toRad(from.longitude - to.longitude);
  const lat1 = toRad(to.latitude);
  const lat2 = toRad(from.latitude);

  const a =
    Math.sin(dLat / 2) * Math.sin(dLat / 2) +
    Math.sin(dLon / 2) * Math.sin(dLon / 2) * Math.cos(lat1) * Math.cos(lat2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
  return R * c;
};

const bearing = (to, from) => {
  const dLon = toRad(from.longitude - to.longitude);
  const lat1 = toRad(to.latitude);
  const lat2 = toRad(from.latitude);
  const y = Math.sin(dLon) * Math.cos(lat2);
  const x =
    Math.cos(lat1) * Math.sin(lat2) -
    Math.sin(lat1) * Math.cos(lat2) * Math.cos(dLon);
  return toDeg(Math.atan2(y, x));
};

class App extends Component {
  constructor() {
    super();
    this.state = {
      heading: 0,
      currentLoc: {
        latitude: 41.276365,
        longitude: -95.990193
      },
      location: null
    };
  }
  toHeading(e) {
    if (e.webkitCompassHeading) {
      return e.webkitCompassHeading;
    }
    return 360 - e.alpha;
  }
  async updateLocation(latitude, longitude) {
    const localTesting = false;
    const baseUrl = localTesting
      ? 'http://localhost:8081/'
      : 'https://us-central1-pizza-compass-232603.cloudfunctions.net/HttpHandler';

    const url = `${baseUrl}?lat=${latitude}&lon=${longitude}`;
    const rawRes = await fetch(url);
    const res = await rawRes.json();
    const [closestPizza] = res;
    this.setState({
      location: {
        name: closestPizza.name,
        latitude: closestPizza.location.lat,
        longitude: closestPizza.location.lng
      }
    });
  }
  componentDidMount() {
    window.addEventListener('deviceorientation', e => {
      this.setState({
        heading: this.toHeading(e),
        alpha: e.alpha,
        wk: e.webkitCompassHeading
      });
    });
    try {
      navigator.geolocation.watchPosition(
        ({ coords }) => {
          const currentLoc = {
            latitude: coords.latitude,
            longitude: coords.longitude
          };
          if (!this.state.location) {
            this.updateLocation(coords.latitude, coords.longitude);
          }
          if (
            currentLoc.latitude !== this.state.currentLoc.latitude &&
            currentLoc.longitude !== this.state.currentLoc.longitude
          ) {
            this.setState({ currentLoc });
          }
        },
        null,
        { enableHighAccuracy: true }
      );
    } catch (e) {
      alert(e);
    }
  }
  render(props, state) {
    return html`
      <${PizzaCompass}
        loc=${state.location}
        heading=${state.heading}
        currentLoc=${state.currentLoc}
      />
    `;
  }
}
const Pizza = ({ rotation }) => {
  console.log(rotation);
  const style = { transform: `rotate(${rotation}deg)` };
  return html`
    <div class="pizza-compass">
      <img style=${style} src="thin-crust-pizza.gif" />
    </div>
  `;
};
const PizzaCompass = ({ loc, heading, currentLoc }) => {
  if (loc && currentLoc) {
    const b = bearing(currentLoc, loc);
    const headingDelta = 180 - (heading - b);
    const distance = Math.floor(distanceFrom(currentLoc, loc) * 10) / 10;
    console.log(distance);
    return html`
      <div class="app">
        <header><h1>${loc.name}</h1></header>
        <${Pizza} rotation=${headingDelta} />
        <footer><h1>${distance} km</h1></footer>
      </div>
    `;
  } else {
    return html`
      <div class="app">
        <header><h1>🍕</h1></header>
        <div class="pizza-compass"></div>
        <footer><h1>🍕</h1></footer>
      </div>
    `;
  }
};

render(
  html`
    <${App} />
  `,
  document.body
);
