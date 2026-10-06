# ADR 0005: Represent Business Dates Without Time-of-Day Semantics

## Status

Accepted

## Context

AgroCore contains business fields that represent calendar dates, such as
agricultural year opening and closing dates.

These values represent a year, month, and day. They do not represent a precise
instant in time and therefore do not require hour, minute, second, or timezone
semantics.

Using `time.Time` directly throughout the domain would expose capabilities that
these business values do not require and could introduce unnecessary timezone
concerns.

The existing agricultural management system also models these business values
as date-only fields.

## Decision

AgroCore will represent business calendar dates using a dedicated `Date` value
object.

A `Date` represents only:

- year;
- month;
- day.

The value object must reject invalid calendar dates.

Its public API must not expose time-of-day or timezone behavior.

The implementation may use Go's `time` package internally when useful, but that
implementation detail must remain encapsulated.

Technical timestamps such as creation and update timestamps are not covered by
this decision and may use `time.Time`.

## Consequences

### Positive

- Business dates have explicit semantics.
- Invalid calendar dates can be rejected at creation time.
- Domain code does not need to handle unnecessary time-of-day information.
- Timezone conversion cannot accidentally change the intended business date.
- The same date representation can be reused by multiple business modules.

### Negative

- Conversion will be required at infrastructure boundaries such as databases,
  JSON, and external APIs.
- AgroCore maintains a small domain type instead of using `time.Time` directly.
- Serialization behavior must be implemented explicitly when required.

## Alternatives Considered

### Use `time.Time` Directly

Rejected for business date-only values because it includes time-of-day and
timezone semantics that are not part of the business concept.

### Store Dates as Strings

Rejected because strings do not guarantee that values represent valid calendar
dates and provide weaker comparison semantics.