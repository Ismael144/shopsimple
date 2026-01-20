from __future__ import annotations

from functools import lru_cache
from typing import Dict
from decimal import Decimal, getcontext, ROUND_HALF_UP
from pathlib import Path 
from typing import Dict
from dataclasses import dataclass
import json 

# High precision for currency math 
getcontext().prec = 28
NANO_FACTOR = Decimal("1e9")
DATA_FILE = Path(__file__).parent / "data" / "currency_conversion.json"

def money_proto_to_decimal(units: int, nanos: int) -> Decimal: 
    """Convert Money proto (units + nanos) to Decimal"""
    return Decimal(units) + Decimal(nanos) / NANO_FACTOR

def decimal_to_money_proto(currency_code: str, amount: Decimal) -> dict: 
    """Convert Decimal amount to Money proto (units + nanos)."""
    # Total nanos 
    total_nanos = (amount * NANO_FACTOR).to_integral_value(rounding=ROUND_HALF_UP)
    units = int(total_nanos // NANO_FACTOR)
    nanos = int(total_nanos % NANO_FACTOR)
    # Handle negative nanos if amount is negative 
    if amount < 0 and nanos > 0: 
        units += 1 
        nanos -= int(NANO_FACTOR)
    return {
        "currency_code": currency_code, 
        "units": units, 
        "nanos": nanos
    }

@lru_cache(maxsize=1)
def get_currency_data() -> Dict[str, Decimal]: 
    """Load currency rates from JSON and cache in memory"""
    try: 
        with open(DATA_FILE, "r", encoding="utf8") as f: 
            raw = json.load(f)
        return {k: Decimal(v) for k, v in raw.items()}
    except Exception as e: 
        raise RuntimeError(f"Failed to load currency data: {e}") from e 

def convert_money(from_money_proto: dict, to_currency: str) -> dict: 
    """Convert money proto to another currency, returns a money proto dict."""
    rates = get_currency_data()
    from_currency = from_money_proto["currency_code"]
    if from_currency not in rates: 
        raise ValueError(f"Unknown currency: {from_currency}")
    if to_currency not in rates: 
        raise ValueError(f"Unknown currency: {to_currency}")
    # Convert proto -> Decimal 
    amount = money_proto_to_decimal(from_money_proto["units"], from_money_proto["nanos"])
    # Convert FROM -> EUR
    eur_amount = amount / rates[from_currency]
    # Convert EUR -> TARGET
    target_amount = eur_amount * rates[to_currency]
    return decimal_to_money_proto(to_currency, target_amount)

def batch_convert_money(from_money_list: list[dict], to_currency: str) -> list[dict]: 
    return [convert_money(m, to_currency) for m in from_money_list]

def get_supported_currencies() -> list[str]: 
    return list(get_currency_data().keys())