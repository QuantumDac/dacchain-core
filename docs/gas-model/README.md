# USD-Fixed Gas Model

## Overview
Web4 DAC Protocol uses a USD-fixed gas model — gas is defined in USD (cents) and converted to DAC using the current DAC price in USD:


gasInDAC = gasInUSD / dacPriceUSD


## Modes
- **manual** — DAC price provided by operator (config/gas.yaml)  
- **oracle** — DAC price retrieved from an external oracle (DEX/CMC) — to be implemented

## Default values
- Normal transaction = **$0.03**
- Bridge transaction = **$0.05**

## Implementation notes
- Exact rational arithmetic used (big.Rat) to avoid floating point rounding errors.
- Result rounded **UP** (ceiling) to 8 decimal places to ensure sufficient fee coverage.
