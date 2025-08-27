# Todo

* **Purpose:** Provides **Linux-specific system properties** that are not meaningful or available on other OSes.
* **Main functions:**

  * `GetProperty(property string, params ...string)`: fetches a property from the **Linux-specific set** (`linuxProperties` map).
  * `linuxProperties` map contains handlers for properties like:


  | Property    | Description                                           |
  |-------------|-------------------------------------------------------|
  | `uuid`      | System UUID from `/sys/class/dmi/id/product_uuid`    |
  | `uname`     | Machine architecture via `uname -m`                  |
  | `osdistro`  | OS distribution (e.g., Ubuntu, Debian)              |
  | `osfamily`  | OS family (e.g., Debian-based, RedHat-based)        |

* **Notes:**

  * Uses Linux system commands (`cat`, `uname`) and `gopsutil`.
  * Only meaningful on **Linux systems**.
  * cf. the same library in the [gocore](https://github.com/abtransitionit/gocore/) project. 

---

a pure GO function to get the OS type:
  - darwin
  - linux
  - windows
    