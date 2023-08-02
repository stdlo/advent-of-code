pub fn process_part1(input: &str) -> String {
    let items = input.lines().map(|v| v.parse::<u32>().unwrap());
    let mut result = 0;
    for a in items.clone() {
        for b in items.clone() {
            if a + b == 2020 {
                result = a * b;
                break;
            }
        }
        if result != 0 {
            break;
        }
    }
    result.to_string()
}

pub fn process_part2(input: &str) -> String {
    let items = input.lines().map(|v| v.parse::<u32>().unwrap());
    let mut result = 0;
    for a in items.clone() {
        for b in items.clone() {
            for c in items.clone() {
                if a + b + c == 2020 {
                    result = a * b * c;
                    break;
                }
            }
        }
        if result != 0 {
            break;
        }
    }
    result.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "1721
979
366
299
675
1456";

    #[test]
    fn part1() {
        let result = process_part1(INPUT);
        assert_eq!(result, "514579");
    }

    #[test]
    fn part2() {
        let result = process_part2(INPUT);
        assert_eq!(result, "241861950");
    }
}
