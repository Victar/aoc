package adventofcode.year2021;

import java.util.ArrayList;

import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day6 extends BaseTest {

	@Test public void runSolver() throws Exception {
		System.out.println(countFishes(80));
	}

	@Test public void runGold() throws Exception {
		System.out.println(countFishes(256));

	}

	public long countFishes(final int round) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day6/input.txt");
		final String[] arr = data.get(0).split(",");
		final Fishes fishes = new Fishes();
		for (final String input : arr) {
			final Integer current = Integer.parseInt(input);
			if (current == 1) {
				fishes.setAge1(fishes.getAge1() + 1);
			}
			if (current == 2) {
				fishes.setAge2(fishes.getAge2() + 1);
			}
			if (current == 3) {
				fishes.setAge3(fishes.getAge3() + 1);
			}
			if (current == 4) {
				fishes.setAge4(fishes.getAge4() + 1);
			}
			if (current == 5) {
				fishes.setAge5(fishes.getAge5() + 1);
			}
			if (current == 6) {
				fishes.setAge6(fishes.getAge6() + 1);
			}
			if (current == 7) {
				fishes.setAge7(fishes.getAge7() + 1);
			}
			if (current == 8) {
				fishes.setAge8(fishes.getAge8() + 1);
			}
		}

		for (int i = 0; i < round; i++) {
			fishes.doRound();
		}
		return fishes.total();
	}

	@Data public static class Fishes {

		long age0;
		long age1;
		long age2;
		long age3;
		long age4;
		long age5;
		long age6;
		long age7;
		long age8;

		public void doRound() {
			final long age0init = this.age0;
			this.age0 = this.age1;
			this.age1 = this.age2;
			this.age2 = this.age3;
			this.age3 = this.age4;
			this.age4 = this.age5;
			this.age5 = this.age6;
			this.age6 = this.age7;
			this.age7 = this.age8;
			this.age6 = this.age6 + age0init;
			this.age8 = age0init;
		}

		public long total() {
			return this.age0 + this.age1 + this.age2 + this.age3 + this.age4 + this.age5 + this.age6 + this.age7 + this.age8;
		}

	}

}
