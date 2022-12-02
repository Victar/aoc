package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Test;

import java.util.ArrayList;

public class Day2 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day2/input.txt");
		int total = 0;
		for (final String input : data) {
			total += new Game(input).getScoreSilver();
		}
		System.out.println(total);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day2/input.txt");
		int total = 0;
		for (final String input : data) {
			total += new Game(input).getScoreGold();
		}
		System.out.println(total);
	}

	@Data class Game {

		Character p1;
		Character p2;

		Game(String input) {
			p1 = input.charAt(0);
			p2 = input.charAt(2);
		}

		//		A,X Rock lose 1
		//		B,Y Paper draw 2
		//		C,Z Scissors win 3
		int getScoreGold() {
			if (p2 == 'X') {
				if (p1 == 'A') {
					return 3 + 0;
				}
				if (p1 == 'B') {
					return 1 + 0;
				}
				if (p1 == 'C') {
					return 2 + 0;
				}
			}
			if (p2 == 'Y') {
				if (p1 == 'A') {
					return 1 + 3;
				}
				if (p1 == 'B') {
					return 2 + 3;
				}
				if (p1 == 'C') {
					return 3 + 3;
				}
			}
			if (p2 == 'Z') {
				if (p1 == 'A') {
					return 2 + 6;
				}
				if (p1 == 'B') {
					return 3 + 6;
				}
				if (p1 == 'C') {
					return 1 + 6;
				}
			}
			return 0;
		}

		int getScoreSilver() {
			if (p1 == 'A') {
				if (p2 == 'X') {
					return 3 + 1;
				}
				if (p2 == 'Y') {
					return 6 + 2;
				}
				if (p2 == 'Z') {
					return 0 + 3;
				}
			}
			if (p1 == 'B') {
				if (p2 == 'X') {
					return 0 + 1;
				}
				if (p2 == 'Y') {
					return 3 + 2;
				}
				if (p2 == 'Z') {
					return 6 + 3;
				}
			}
			if (p1 == 'C') {
				if (p2 == 'X') {
					return 6 + 1;
				}
				if (p2 == 'Y') {
					return 0 + 2;
				}
				if (p2 == 'Z') {
					return 3 + 3;
				}
			}
			return 0;
		}

	}

}
