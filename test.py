from datetime import date

def get_end_of_month(year: int, month:int) -> date:
    if month == 12:
        return date(year + 1, 1, 1) - date.resolution
    return date(year, month + 1, 1) - date.resolution

from enum import Enum

class State(Enum):
    BEFORE = 1
    SCHEDULED = 2
    IN_PROGRESS = 3
    COMPLETED = 4
    CANCELED = 5
    FAILED = 6