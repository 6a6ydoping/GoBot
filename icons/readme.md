# Icons Directory

This directory contains a collection of icons used in the project. Each icon represents different weather conditions and is associated with a specific code.

## Icon List

- `01d@2x.png`: Clear sky (day)
- `01n@2x.png`: Clear sky (night)
- `02d@2x.png`: Few clouds (day)
- `02n@2x.png`: Few clouds (night)
- `03d@2x.png`: Scattered clouds (day)
- `03n@2x.png`: Scattered clouds (night)
- `04d@2x.png`: Broken clouds (day)
- `04n@2x.png`: Broken clouds (night)
- `09d@2x.png`: Shower rain (day)
- `09n@2x.png`: Shower rain (night)
- `10d@2x.png`: Rain (day)
- `10n@2x.png`: Rain (night)
- `11d@2x.png`: Thunderstorm (day)
- `11n@2x.png`: Thunderstorm (night)
- `13d@2x.png`: Snow (day)
- `13n@2x.png`: Snow (night)
- `50d@2x.png`: Mist (day)
- `50n@2x.png`: Mist (night)

## Why Manual Download?

Although it might be tempting to make requests to a remote server for each icon, the decision to manually download and store the icons locally on the server is made for the following reasons:

1. **Reduced Latency:** By storing the icons locally, you reduce the latency associated with making external requests. This is especially beneficial for frequently accessed assets like weather icons.

2. **Reliability:** Local storage ensures that the icons are always available, even if the external server experiences downtime or changes its structure.

3. **Bandwidth Efficiency:** Storing icons locally reduces the need for repetitive downloads, saving bandwidth for both your server and the external server providing the icons.


Keep in mind that this approach is suitable when the number of icons is relatively small, and the benefits of local storage outweigh the potential downsides. (I've included all the icons in the repository, even though storing binary files in a repository is generally considered a bad practice. I opted for this approach to enhance the project's clarity and simplify comprehension.)
